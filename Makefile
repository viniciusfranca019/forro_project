serve:
	@docker-compose -f .docker/prd/docker-compose.yml up
build-and-serve:
	@docker-compose -f .docker/prd/docker-compose.yml up --build