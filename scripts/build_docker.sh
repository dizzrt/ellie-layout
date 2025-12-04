#!/bin/sh

# exit when any command fails
set -e

cd "$(dirname "$0")/../"
. ./scripts/build_init.sh

echo "Start building docker image ellie-layout:${_VERSION}..."

docker buildx build -f ./scripts/Dockerfile \
    --platform="${PLATFORM}" \
    --build-arg GOOS="${GOOS}" \
    --build-arg GOARCH="${GOARCH}" \
    --build-arg ENV="${ENV}" \
    --build-arg VERSION="${_VERSION}" \
    --build-arg GIT_COMMIT="${GIT_COMMIT}" \
    -t ellie-layout/ellie-layout:${_VERSION} .

echo "${GREEN}Completed building docker image ${_VERSION}.${NC}"
echo ""
echo "Command to start ellie-layout"
echo ""
echo "$ docker run -d -p 8081:8081 -p 50051:50051 -v ~/.ellie-layout/logs:/app/logs --name ellie-layout ellie-layout/ellie-layout:${_VERSION}"
echo ""
