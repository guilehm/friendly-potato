DOCKER_COMPOSE=docker-compose

build:
	-$(DOCKER_COMPOSE) build

setup:
	- cp docker-compose.sample.yml docker-compose.yml

run:
	-$(DOCKER_COMPOSE) up
