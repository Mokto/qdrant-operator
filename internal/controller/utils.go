package controller

import (
	"crypto/sha256"
	"fmt"
	"sort"
)

func hashMap(m map[string]string) string {
	h := sha256.New()

	// hashing maps is not deterministic because it is not sorted, so we have
	// to manually hash each key value into another hash after sorting

	// first we hash the string representation of the bool
	// h.Write([]byte(fmt.Sprintf("%v", b)))
	keys := make([]string, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := m[k]
		// hash the key first then value before writing to the final sum
		// this is because if we just wrote the key and value as strings,
		// the following two maps would be equivalent:
		// { "k1" : "v1" } and { "k" : "1v1" }
		// so if we first hash the inputs, that ensures the final hash is
		// always unique unless there is a sha256 hash collision
		b := sha256.Sum256([]byte(fmt.Sprintf("%v", k)))
		h.Write(b[:])
		b = sha256.Sum256([]byte(fmt.Sprintf("%v", v)))
		h.Write(b[:])
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
