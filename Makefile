help:
	@echo "-------------------- Comandos Makefile ------------------"
	@echo "make start-dev =====>> inicia o ambiente de desenvolvimento"

start-dev:
	docker compose -f build/dev/docker-compose.yml up