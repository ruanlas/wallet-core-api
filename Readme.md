[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/) ![Badge](https://img.shields.io/static/v1?label=Go&message=1.21.3&color=00ADD8&style=flat&logo=Go) ![Badge](https://img.shields.io/static/v1?label=MySQL&message=8.2.0&color=4479A1&style=flat&logo=MySQL) [![Github tag](https://badgen.net/github/tag/Naereen/Strapdown.js)](https://github.com/ruanlas/wallet-core-api/tags/)

# wallet-core-api

O objetivo deste projeto é oferecer recursos para a gestão financeira pessoal.

Ele permitirá lançar projeções de receitas e gastos, além de permitir lançar também os gastos e receitas efetivadas.

## Índice
<!--ts-->
   * [Tecnologias utilizadas neste projeto](#tecnologias-utilizadas-neste-projeto)
   * [Pré-requisitos](#pré-requisitos)
   * [Ambiente de desenvolvimento](#ambiente-de-desenvolvimento)
      * [Comandos úteis](#comandos-úteis)
   * [Documentação](#documentação)
   * [Monitoramento](#monitoramento)
   * [Variáveis de ambiente](#variáveis-de-ambiente)
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
| ELASTIC_APM_SERVICE_NAME  | Nome do serviço no APM  |
| ELASTIC_APM_SERVICE_VERSION  | Versão do serviço no APM  |
| ELASTIC_APM_SERVER_URL  | Host/URL do serviço do APM  |

