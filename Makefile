serve:
	@docker-compose -f .docker/prd/docker-compose.yml up
build-and-serve:
	@docker-compose -f .docker/prd/docker-compose.yml up --build
seeds:
	@docker-compose -f .docker/prd/docker-compose.yml exec go_cli "./main -run seeds"