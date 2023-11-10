help:
	@echo "-------------------------------------------------------------------------------------"
	@echo "-------------------------------- COMANDOS DO PROJETO --------------------------------"
	@echo "-------------------------------------------------------------------------------------"
	@echo "- Ambiente de desenvolvimento -------------------------------------------------------"
	@echo "# make dev-start       =====>> inicia o ambiente de desenvolvimento"
	@echo "# make dev-stop        =====>> interrompe o ambiente de desenvolvimento"
	@echo "# make dev-init-load   =====>> carrega as informações iniciais no banco de dados"
	@echo "# make dev-drop-tables =====>> remove todos os registros do banco de dados"
	@echo "-------------------------------------------------------------------------------------"

dev-start:
	docker compose -f build/dev/docker-compose.yml up -d

dev-stop:
	docker compose -f build/dev/docker-compose.yml down

dev-init-load:
	docker exec -i wallet-db sh -c 'exec mysql -uroot -p123456 --default-character-set=utf8mb4 wallet_core < /docker-entrypoint-initdb.d/database_init_load.sql'

dev-drop-tables:
	docker exec -i wallet-db sh -c 'exec mysql -uroot -p123456 --default-character-set=utf8mb4 wallet_core < /docker-entrypoint-initdb.d/database_drop_tables.sql'