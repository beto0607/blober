BINARY_NAME=blober

build:
	go build -o ${BINARY_NAME} main.go

docker-build-push:
	docker buildx build --platform linux/amd64,linux/arm64 -t docker-registry.ralb.dev/blober:latest --push .

