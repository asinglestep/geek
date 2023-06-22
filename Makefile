export DOCKER_TAG=$(shell git rev-parse --short HEAD)

generate:
	swag init --dir ./httpserver --output ./httpserver/docs

httpserver-image:
	docker build -f httpserver/build/Dockerfile . -t winter2023/httpserver:$(DOCKER_TAG)

