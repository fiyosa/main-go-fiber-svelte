#!/bin/sh
set -e

echo "Deploying..."
docker build -t portfolio -f build/package/Dockerfile .
docker stop portfolio 2>/dev/null || true
docker rm portfolio 2>/dev/null || true

docker run -itd \
  --name portfolio \
  --restart always \
  -p 8000:8000 \
  --network proxy \
  --env-file /root/docker/go-fiber-svelte/.env \
  portfolio

docker image prune -f

printf "\nDeploy success"