# Write a docker for nodejs
FROM oven/bun:1
WORKDIR /app
COPY package.json /app
COPY bun.lockb /app
RUN bun install
COPY . /app
CMD ["bun", "run.js"]
