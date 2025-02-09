SHORT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo dev)

.PHONY: build
build:
	go build -o blb ./cmd/... 

.PHONY: docker-build
docker-build:
	docker build -t slikshooz/bambulab-bridge:latest -f ./docker/Dockerfile .
	docker tag slikshooz/bambulab-bridge:latest slikshooz/bambulab-bridge:$(SHORT_COMMIT)

.PHONY: docker-push
docker-push: docker-build
	docker push slikshooz/bambulab-bridge:latest
	docker push slikshooz/bambulab-bridge:$(SHORT_COMMIT)