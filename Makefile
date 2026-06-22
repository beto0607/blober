BINARY_NAME=blober
IMAGE_NAME=blober

build:
	go build -o ${BINARY_NAME} main.go

docker-build:
	docker buildx build --platform linux/amd64,linux/arm64 -t docker-registry.ralb.dev/${IMAGE_NAME}:latest .

docker-build-push:
	docker buildx build --platform linux/amd64,linux/arm64 -t docker-registry.ralb.dev/${IMAGE_NAME}:latest --push .

