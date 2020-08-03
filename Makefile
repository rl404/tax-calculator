.DEFAULT_GOAL := docker

# Docker-compose.yml path.
COMPOSE_PATH := ./deployment/docker-compose.yml

# Standarize go coding style for the whole app.
.PHONY: fmt
fmt:
	@go fmt ./...

# Clean project binary,
.PHONY: clean
clean:
	@go clean ./...

# Build docker images and containers.
.PHONY: docker-build
docker-build: clean fmt
	@docker-compose -f $(COMPOSE_PATH) -p tax-calculator build
	@docker image prune -f --filter label=stage="tc_builder"

# Start built docker images and containers.
,PHONY: docker-up
docker-up:
	@docker-compose -f $(COMPOSE_PATH) -p tax-calculator up -d
	@docker logs --follow tax-calculator

# Build and start docker images and containers.
.PHONY: docker
docker: docker-build docker-up

# Stop all related docker containers.
.PHONY: docker-stop
docker-stop:
	@docker stop tax-calculator tc_db

# Stop and remove all related docker container.
.PHONY: docker-rm
docker-rm:
	@docker rm tax-calculator tc_db || echo ""
	@docker volume rm tax-calculator_postgres-volume || echo ""
	@docker rmi taxcalculator
