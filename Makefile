SHELL=/bin/bash
.DEFAULT_GOAL := help

.PHONY: help
help:
	@echo -e "Welcome to Go-Intro. \nPlease use \`make <target>\` where <target> is one of the following:"
	@echo "  all    run both \`make up\` and \`make exec\` commands"
	@echo "  up     start the docker service(s)"
	@echo "  exec   execute the bash command within the go-intro service"
	@echo "  down   stop and remove the running docker service(s)"
	@echo "  help   print this help message"

.PHONY: all
all: up exec

.PHONY: up
up:
	docker compose up -d

.PHONY: exec
exec:
	docker exec -it go-intro bash

.PHONY: down
down:
	docker compose down -v
