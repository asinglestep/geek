export DOCKER_TAG=$(shell git rev-parse --short HEAD)

httpserver-image:
	docker build -f httpserver/build/Dockerfile . -t httpserver:$(DOCKER_TAG)