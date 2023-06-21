export DOCKER_TAG=$(shell git rev-parse --short HEAD)

generate:
	swag init --dir ./httpserver

httpserver-image:
	docker build -f httpserver/build/Dockerfile . -t httpserver:$(DOCKER_TAG)

