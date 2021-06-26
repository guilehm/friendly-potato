DOCKER_COMPOSE=docker-compose


setup:
	- cp docker-compose.sample.yml docker-compose.yml

run:
	-$(DOCKER_COMPOSE) --env-file .env up
