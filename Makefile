SHELL=/bin/bash
.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo -e "Welcome to Go-Intro. \nPlease use \`make <target>\` where <target> is one of the following:"
	@echo "  up     start the docker service(s)"
	@echo "  down   stop and remove the running docker service(s)"
	@echo "  help   print this help message"

.PHONY: up
up:
	docker compose up -d

.PHONY: down
down:
	docker compose down -v
