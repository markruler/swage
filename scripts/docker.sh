#!/usr/bin/env bash

set -e

BINARY="swage"
IMAGE="swage"
REPO="markruler"
REPO_PORT=""
VERSION=$(cat ./VERSION)

remove::docker_images() {
	image_hash=$(docker images | grep $IMAGE | awk '{print $3}')
	if [ -z "$image_hash" ]; then
		echo "\$image_hash is NULL"
	else
		echo "\$image_hash is NOT NULL"
		docker rmi -f $image_hash
	fi
}

echo "Clean..."
remove::docker_images
mkae clean

echo "Build $BINARY..."
make build
echo "Build Complete!"

echo "MAke a image..."
docker build \
--tag $IMAGE:$VERSION \
--tag $IMAGE:latest \
--build-arg VERSION=$VERSION \
--file ./Dockerfile \
$PWD

echo "Tag..."
docker tag $IMAGE:$VERSION $REPO$REPO_PORT/$IMAGE:$VERSION
docker tag $IMAGE:$VERSION $REPO$REPO_PORT/$IMAGE:latest

# echo "Push..."
# docker push $REPO$REPO_PORT/$IMAGE:$VERSION
# docker push $REPO$REPO_PORT/$IMAGE:latest

echo -e "\n>>> PRINT IMAGES <<<"
echo "$(docker images | grep $IMAGE | awk '{ printf ("%s\n", $0) }')"

echo "Run..."
docker run \
--rm \
--interactive \
--tty \
--volume $PWD/examples/testdata:/testdata \
swage:$VERSION gen /testdata/json/editor.swagger.json \
--output /testdata/docker-swage.xlsx \
--verbose
