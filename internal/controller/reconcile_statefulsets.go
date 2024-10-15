package controller

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	v1 "k8s.io/api/apps/v1"
	v1core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"

	qdrantv1alpha1 "qdrantoperator.io/operator/api/v1alpha1"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (r *QdrantClusterReconciler) reconcileStatefulsets(ctx context.Context, log logr.Logger, obj *qdrantv1alpha1.QdrantCluster, checksum string) error {

	cordonedPeerIds := []string{}

	falseValue := false
	trueValue := true
	securityUser := int64(1000)
	securityGroup := int64(2000)

	leader := obj.Status.Peers.GetLeader()
	indexStatefulSet := 0
	for _, statefulSetConfig := range obj.Spec.Statefulsets {
		if indexStatefulSet > 0 && leader == nil {
			continue
		}
		labels := map[string]string{
			"cluster": obj.Name,
			"name":    obj.Name + "-" + statefulSetConfig.Name,
		}
		if statefulSetConfig.EphemeralStorage {
			labels["qdrant-ephemeral-storage"] = "true"
		} else {
			labels["qdrant-ephemeral-storage"] = "false"
		}
		var volumeClaimTemplates []v1core.PersistentVolumeClaim
		volumes := []v1core.Volume{{
			Name: "qdrant-init",
			VolumeSource: v1core.VolumeSource{
				EmptyDir: &v1core.EmptyDirVolumeSource{},
			},
		}, {
			Name: "qdrant-snapshots",
			VolumeSource: v1core.VolumeSource{
				EmptyDir: &v1core.EmptyDirVolumeSource{},
			},
		}, {
			Name: "qdrant-config",
			VolumeSource: v1core.VolumeSource{
				ConfigMap: &v1core.ConfigMapVolumeSource{
					LocalObjectReference: v1core.LocalObjectReference{
						Name: obj.Name,
					},
					DefaultMode: &[]int32{493}[0],
				},
			},
		}}
		if statefulSetConfig.VolumeClaim != nil {
			volumeClaimTemplates = []v1core.PersistentVolumeClaim{{
				ObjectMeta: v1meta.ObjectMeta{Name: "qdrant-storage"},
				Spec: v1core.PersistentVolumeClaimSpec{
					AccessModes:      []v1core.PersistentVolumeAccessMode{v1core.ReadWriteOnce},
					StorageClassName: &statefulSetConfig.VolumeClaim.StorageClassName,
					Resources: v1core.VolumeResourceRequirements{
						Requests: v1core.ResourceList{
							"storage": resource.MustParse(statefulSetConfig.VolumeClaim.Storage),
						},
					},
				},
			}}
		} else {
			volumes = append(volumes, v1core.Volume{
				Name: "qdrant-storage",
				VolumeSource: v1core.VolumeSource{
					EmptyDir: &v1core.EmptyDirVolumeSource{},
				},
			})
		}
		affinity := statefulSetConfig.Affinity
		if affinity == nil {
			affinity = &v1core.Affinity{}
		}
		if affinity.PodAntiAffinity == nil {
			affinity.PodAntiAffinity = &v1core.PodAntiAffinity{}
		}
		affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution = []v1core.PodAffinityTerm{{
			LabelSelector: &v1meta.LabelSelector{MatchLabels: map[string]string{"cluster": obj.Name}},
			TopologyKey:   "kubernetes.io/hostname",
		}}
		containers := []v1core.Container{{
			Name:  "qdrant",
			Image: obj.Spec.Image,
			Command: []string{
				"/bin/bash",
				"-c",
			},
			Args: []string{
				"./config/initialize.sh",
			},
			Env: []v1core.EnvVar{{
				Name:  "QDRANT_INIT_FILE_PATH",
				Value: "/qdrant/init/.qdrant-initialized",
			}},
			Resources: statefulSetConfig.Resources,
			Lifecycle: &v1core.Lifecycle{
				PreStop: &v1core.LifecycleHandler{
					Exec: &v1core.ExecAction{
						Command: []string{"sleep", "3"},
					},
				},
			},
			Ports: []v1core.ContainerPort{{
				ContainerPort: 6333,
				Name:          "http",
				Protocol:      v1core.Protocol("TCP"),
			}, {
				ContainerPort: 6334,
				Name:          "grpc",
				Protocol:      v1core.Protocol("TCP"),
			}, {
				ContainerPort: 6335,
				Name:          "p2p",
				Protocol:      v1core.Protocol("TCP"),
			}},
			ReadinessProbe: &v1core.Probe{
				FailureThreshold: 6,
				ProbeHandler: v1core.ProbeHandler{HTTPGet: &v1core.HTTPGetAction{Path: "/readyz", Port: intstr.IntOrString{
					IntVal: 6333,
				}, Scheme: "HTTP"}},
				InitialDelaySeconds: 5,
				PeriodSeconds:       5,
				SuccessThreshold:    1,
				TimeoutSeconds:      1,
			},
			SecurityContext: &v1core.SecurityContext{
				AllowPrivilegeEscalation: &falseValue,
				Privileged:               &falseValue,
				ReadOnlyRootFilesystem:   &trueValue,
				RunAsUser:                &securityUser,
				RunAsGroup:               &securityGroup,
				RunAsNonRoot:             &trueValue,
			},
			VolumeMounts: []v1core.VolumeMount{{
				MountPath: "/qdrant/storage",
				Name:      "qdrant-storage",
			}, {
				MountPath: "/qdrant/config/initialize.sh",
				Name:      "qdrant-config",
				SubPath:   "initialize.sh",
			}, {
				MountPath: "/qdrant/config/production.yaml",
				Name:      "qdrant-config",
				SubPath:   "production.yaml",
			}, {
				MountPath: "/qdrant/snapshots",
				Name:      "qdrant-snapshots",
			}, {
				MountPath: "/qdrant/init",
				Name:      "qdrant-init",
			}},
		}}
		if statefulSetConfig.EphemeralStorage {
			containers = append(containers, v1core.Container{
				Name: "ready-container-check",
				// Image: "ghcr.io/mokto/qdrant-ready-container-check:0.10.0",
				Image: "ghcr.io/mokto/qdrant-ready-container-check:0.11.0",
				Command: []string{
					"sleep",
					"315000000", // 10 years
				},
				Env: []v1core.EnvVar{{
					Name:  "API_KEY",
					Value: obj.Spec.ApiKey,
				}},
				StartupProbe: &v1core.Probe{
					ProbeHandler: v1core.ProbeHandler{
						Exec: &v1core.ExecAction{
							Command: []string{
								"bun", "run.js",
							},
						},
					},
					InitialDelaySeconds: 30,
					PeriodSeconds:       5,
					SuccessThreshold:    1,
					FailureThreshold:    10 * 3600,
					TimeoutSeconds:      1,
				},
			})
		}
		podTemplate := v1core.PodTemplateSpec{
			ObjectMeta: v1meta.ObjectMeta{Labels: labels, Annotations: map[string]string{"checksum": checksum}},
			Spec: v1core.PodSpec{
				InitContainers: []v1core.Container{{
					Name:  "ensure-dir-ownership",
					Image: obj.Spec.Image,
					Command: []string{
						"chown",
						"-R",
						"1000:3000",
						"/qdrant/storage",
						"/qdrant/snapshots",
					},
					VolumeMounts: []v1core.VolumeMount{{
						MountPath: "/qdrant/storage",
						Name:      "qdrant-storage",
					}, {
						MountPath: "/qdrant/snapshots",
						Name:      "qdrant-snapshots",
					}},
				}},
				Containers:        containers,
				PriorityClassName: statefulSetConfig.PriorityClassName,
				NodeSelector:      statefulSetConfig.NodeSelector,
				ImagePullSecrets:  obj.Spec.ImagePullSecrets,
				Affinity:          affinity,
				Tolerations:       statefulSetConfig.Tolerations,
				Volumes:           volumes,
			},
		}
		statefulSet := &v1.StatefulSet{}
		if err := r.Get(ctx, types.NamespacedName{
			Name:      obj.Name + "-" + statefulSetConfig.Name,
			Namespace: obj.Namespace,
		}, statefulSet); err != nil {
			replicas := &statefulSetConfig.Replicas
			if leader == nil {
				replicas = new(int32)
				*replicas = 1
			}
			statefulSet = &v1.StatefulSet{
				ObjectMeta: v1meta.ObjectMeta{
					Name:      obj.Name + "-" + statefulSetConfig.Name,
					Namespace: obj.Namespace,
					OwnerReferences: []v1meta.OwnerReference{{
						APIVersion: obj.APIVersion,
						Kind:       obj.Kind,
						Name:       obj.Name,
						UID:        obj.UID,
					}},
				},
				Spec: v1.StatefulSetSpec{
					PersistentVolumeClaimRetentionPolicy: &v1.StatefulSetPersistentVolumeClaimRetentionPolicy{
						WhenDeleted: "Delete",
						WhenScaled:  "Delete",
					},
					Replicas:            replicas,
					Selector:            &v1meta.LabelSelector{MatchLabels: labels},
					PodManagementPolicy: v1.PodManagementPolicyType("Parallel"),
					// PodManagementPolicy: v1.PodManagementPolicyType("OrderedReady"),
					ServiceName: obj.GetHeadlessServiceName(),

					Template:             podTemplate,
					VolumeClaimTemplates: volumeClaimTemplates,
				},
			}

			if err := r.Client.Create(ctx, statefulSet); err != nil {
				return err
			}
		} else {

			if leader == nil {
				replicas := int32(1)
				statefulSet.Spec.Replicas = &replicas
			} else {
				// if the number decreases, we need to scale down slowly & cordon first
				if statefulSetConfig.Replicas < *statefulSet.Spec.Replicas {
					peerId := obj.Status.Peers.FindPeerId(fmt.Sprintf("%s-%d", statefulSetConfig.Name, *statefulSet.Spec.Replicas-1))
					isPeerEmpty := true
					for _, collection := range obj.Status.Collections {
						for peerIdWithShards := range collection.Shards {
							if peerIdWithShards == peerId {
								isPeerEmpty = false
								break
							}
						}
					}
					if isPeerEmpty {
						wantedReplicas := *statefulSet.Spec.Replicas - 1
						statefulSet.Spec.Replicas = &wantedReplicas
					} else {
						log.Info("We should scale down slowly just by marking the last node for deletion")
						cordonedPeerIds = append(cordonedPeerIds, peerId)
					}

				} else if statefulSetConfig.Replicas > *statefulSet.Spec.Replicas {
					log.Info("We can scale immediately")
					statefulSet.Spec.Replicas = &statefulSetConfig.Replicas
				}
			}
			statefulSet.Spec.Template = podTemplate

			if err := r.Client.Update(ctx, statefulSet); err != nil {
				return err
			}
		}
		indexStatefulSet++
	}

	patch := client.MergeFrom(obj.DeepCopy())
	obj.Status.CordonedPeerIds = cordonedPeerIds

	err := r.Client.Status().Patch(ctx, obj, patch)
	if err != nil {
		log.Error(err, "unable to update QdrantCluster status")
	}

	return nil
}
