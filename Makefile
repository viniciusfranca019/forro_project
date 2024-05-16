serve:
	@docker-compose -f .docker/prd/docker-compose.yml up
build-and-serve:
	@docker-compose -f .docker/prd/docker-compose.yml up --build
dev:
	@go build -o ./tmp/main ./cmd && air
dev-setup:
	@docker-compose -f .docker/dev/docker-compose.yml up --build