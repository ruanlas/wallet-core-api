help:
	@echo "-----------------------------------------------------------------------------------------------"
	@echo "-------------------------------------- VARIÁVEIS DO MAKE --------------------------------------"
	@echo "-----------------------------------------------------------------------------------------------"
	@echo "# IMAGE_NAME           =====>> Nome da imagem docker"
	@echo "# TAG                  =====>> Tag/versão da imagem docker"
	@echo "-----------------------------------------------------------------------------------------------"
	@echo "------------------------------------- COMANDOS DO PROJETO -------------------------------------"
	@echo "-----------------------------------------------------------------------------------------------"
	@echo "- Ambiente de desenvolvimento -----------------------------------------------------------------"
	@echo "# make dev-start         =====>> Inicia o ambiente de desenvolvimento"
	@echo "# make dev-stop          =====>> Interrompe o ambiente de desenvolvimento"
	@echo "# make dev-init-load     =====>> Carrega as informações iniciais no banco de dados"
	@echo "# make dev-drop-tables   =====>> Remove todos os registros do banco de dados"
	@echo "# make dev-datafake-load =====>> Carrega o banco de dados com dados fakes"
	@echo "-----------------------------------------------------------------------------------------------"
	@echo "- Comandos da aplicação -----------------------------------------------------------------------"
	@echo "# make test            =====>> Executa os testes unitários do projeto"
	@echo "# make build-go-app    =====>> Faz o build da aplicação e gera o binário em ./build/app/"
	@echo "# make build-image     =====>> Gera a imagem docker do projeto"
	@echo "# make push-image      =====>> Envia a imagem docker para o repositório de imagens"
	@echo "# make doc-generate    =====>> Gera a documentação em swagger e também converte para markdown"
	@echo "-----------------------------------------------------------------------------------------------"

dev-start:
	docker compose -f build/dev/docker-compose.yml up -d

dev-stop:
	docker compose -f build/dev/docker-compose.yml down

dev-init-load:
	docker exec -i wallet-db sh -c 'exec mysql -uroot -p123456 --default-character-set=utf8mb4 wallet_core < /docker-entrypoint-initdb.d/database_init_load.sql'

dev-datafake-load:
	docker exec -i wallet-db sh -c 'exec mysql -uroot -p123456 --default-character-set=utf8mb4 wallet_core < /docker-entrypoint-initdb.d/database_datafakes.sql'

dev-drop-tables:
	docker exec -i wallet-db sh -c 'exec mysql -uroot -p123456 --default-character-set=utf8mb4 wallet_core < /docker-entrypoint-initdb.d/database_drop_tables.sql'

test:
	mkdir -p ./build/test-result
	go test -short -coverprofile=build/test-result/cov.out `go list ./... | grep -v vendor`

build-go-app:
	$(MAKE) -C ./scripts/build/ build

build-image: build-go-app
ifndef IMAGE_NAME
	$(error IMAGE_NAME is not set)
endif
ifndef TAG
	$(error TAG is not set)
endif
	@echo "docker build . -t ruanlas/$(IMAGE_NAME):$(TAG)"

push-image: build-image
	@echo "docker push ruanlas/$(IMAGE_NAME):$(TAG)"

doc-generate:
	@echo "Gerando documentação do swagger"
	docker run --rm -it -v "$(PWD):/work" ruanlas/go-swagger-generator:v1.0.0 swag init -g cmd/main.go
	@echo "Convertendo a documentação do swagger em markdown"
	docker run --rm -it -v "$(PWD):/work" ruanlas/swagger-to-markdown-convert:v1.0.0 swagger-markdown -i ./docs/swagger.yaml -o ./api_doc.md
	@echo "Concluído"
