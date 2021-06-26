DOCKER_COMPOSE=docker-compose


run:
	-$(DOCKER_COMPOSE) --env-file .env up
