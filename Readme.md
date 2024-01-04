[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/) ![Badge](https://img.shields.io/static/v1?label=Go&message=1.21.3&color=00ADD8&style=flat&logo=Go) ![Badge](https://img.shields.io/static/v1?label=MySQL&message=8.2.0&color=4479A1&style=flat&logo=MySQL) [![Github tag](https://badgen.net/github/tag/Naereen/Strapdown.js)](https://github.com/ruanlas/wallet-core-api/tags/)

# wallet-core-api

O objetivo deste projeto é oferecer uma api com recursos para a gestão financeira pessoal.

Os recursos disponíveis no momento são:
   * Cadastro de projeção de receitas
   * Cadastro de receitas
   * Cadastro de projeção de despesas
   * Cadastro de despesas

## Índice
<!--ts-->
   * [Tecnologias utilizadas neste projeto](#tecnologias-utilizadas-neste-projeto)
   * [Pré-requisitos](#pré-requisitos)
   * [Ambiente de desenvolvimento](#ambiente-de-desenvolvimento)
      * [Comandos úteis](#comandos-úteis)
      * [Obtendo o token do usuário](#obtendo-o-token-do-usuário)
         * [Autenticando na API do keycloak](#autenticando-na-api-do-keycloak)
   * [Documentação](#documentação)
   * [Monitoramento](#monitoramento)
   * [Variáveis de ambiente](#variáveis-de-ambiente)
   * [Documentação de referência](#documentação-de-referência)
<!--te-->

## Tecnologias utilizadas neste projeto
 - Go 1.21.3
 - MySQL 8.2.0
 - APM Server 7.17.14
 - Keycloak 22.0.5

## Pré-requisitos

Para rodar o projeto localmente, é necessário ter instalado as seguintes ferramentas:
- Make versão 4.3 (ou mais recente)
- Docker versão 24.0.5 (ou mais recente)
- Docker Compose versão 2.20.2 (ou mais recente)
- Go versão 1.21.3 (ou mais recente)

## Ambiente de desenvolvimento

Para iniciar o ambiente de desenvolvimento, basta executar o comando `make dev-start` e todas as dependências serão carregadas e inicializadas.

Para visualizar os comandos do projeto basta executar `make help` e aparecerá listado todos os comandos disponíveis.

### Comandos úteis
Inicia as dependências e sobe todos os containers do ambiente de desenvolvimento:
```bash
$ make dev-start
```
Pausa todos os containers do ambiente de desenvolvimento:
```bash
$ make dev-stop
```
Faz a carga inicial dos dados de domínio do banco de dados:
```bash
$ make dev-init-load
```
Limpa todos os registros do banco de dados:
```bash
$ make dev-drop-tables
```
Carrega o banco de dados com dados fakes:
```bash
$ make dev-datafake-load
```

### Obtendo o token do usuário
Para utilizar esta API é necessário possuir um token válido. O token pode ser obtido de duas formas:
 1) Fazendo uma autenticação por oauth2 no keycloak e obter o token
 2) Se autenticar na API do keycloak e obter o token válido

#### Autenticando na API do keycloak
Para autenticar na API do keycloak, é necessário fazer uma chamada `POST` no endpoint `/realms/{realm}/protocol/openid-connect/token`.
Segue abaixo um exemplo de chamada por `curl`:
```bash
$ curl --location '{keycloak-host}/realms/wallet/protocol/openid-connect/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=password' \
--data-urlencode 'client_id={client_id}' \
--data-urlencode 'client_secret={client_secret}' \
--data-urlencode 'username={username}' \
--data-urlencode 'password={password}' \
--data-urlencode 'scope=openid'
```
Para atualizar o token basta fazer outra chamada para o mesmo endpoint. Segue abaixo um exemplo de chamada por `curl` para atualizar o token:
```bash
$ curl --location '{keycloak-host}/realms/master/protocol/openid-connect/token' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=refresh_token' \
--data-urlencode 'client_id={client_id}' \
--data-urlencode 'client_secret={client_secret}' \
--data-urlencode 'refresh_token={refresh_token}'
```

## Documentação
A documentação da API deste projeto foi feita utilizando o projeto [swaggo/swag](https://github.com/swaggo/swag), que faz a conversão de anotações em Go para documentação Swagger 2.0. 
Para acessar a documentação em Swagger, basta executar o projeto e acessar o endereço `http://localhost:8080/swagger/index.html` no navegador. 
Caso queira acessar a documentação pelo repositório, existe também uma documentação em markdown localizaca em [api_doc.md](api_doc.md).

**OBS:** A documentação é gerada automaticamente no Github por uma action toda vez que um código é mergeado ou enviado para a branch `main`. No entanto, é possível também gerar a documentação automaticamente em ambiente de desenvolvimento (caso seja necessário) utilizando o comando abaixo:
```bash
$ make doc-generate
```

## Monitoramento
O projeto está utilizando o Prometheus, e as métricas de monitoramento está disponível no endpoint `http://localhost:2112/metrics`, e o estado da aplicação está disponível no endpoint `http://localhost:2112/metrics`. A porta que o Prometheus estará sendo executado é definido pela variável de ambiente `PROMETHEUS_PORT`.

## Variáveis de ambiente
|  Nome |  Descrição |
|---|---|
| SERVICE_HOST  | Host/URL de acesso ao serviço  |
| SERVICE_PORT  | Porta que o serviço será executado |
| PROMETHEUS_PORT  | Porta que o prometheus estará executando  |
| DATABASE_HOST  | Host do banco de dados  |
| DATABASE_PORT  | Porta do banco de dados  |
| DATABASE_NAME  | Nome do banco de dados  |
| DATABASE_USERNAME  | Usuário do banco de dados  |
| DATABASE_PASSWORD  | Senha do usuário do banco de dados  |
| IDP_HOST  | Host do keycloak  |
| IDP_PORT  | Porta do keycloak  |
| IDP_MAIN_REALM  | Realm principal do keycloak. Geralmente o Realm `master`  |
| IDP_USER_ADMIN  | Usuário administrador do keycloak  |
| IDP_PASSWORD_ADMIN  | Senha do usuário administrador do keycloak  |
| IDP_REALM  | Realm do keycloak que será usado para a aplicação  |
| IDP_CLIENT_IDENTIFIER  | Id do client da API do keycloak  |
| IDP_CLIENT_SECRET  | Secret do client da API do keycloak  |
| ELASTIC_APM_SERVICE_NAME  | Nome do serviço no APM  |
| ELASTIC_APM_SERVICE_VERSION  | Versão do serviço no APM  |
| ELASTIC_APM_SERVER_URL  | Host/URL do serviço do APM  |

## Documentação de referência
[Keycloak](https://www.keycloak.org/documentation)

[OpenID Connect Keycloak](https://www.keycloak.org/docs/latest/securing_apps/#_oidc)

[Keycloak API](https://www.keycloak.org/docs-api/23.0.3/rest-api/index.html)

[Gocloak - Golang Keycloak API Package](https://github.com/Nerzal/gocloak)
