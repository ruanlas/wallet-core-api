![Badge](https://img.shields.io/static/v1?label=Go&message=1.21.3&color=00ADD8&style=flat&logo=Go) ![Badge](https://img.shields.io/static/v1?label=MySQL&message=8.2.0&color=4479A1&style=flat&logo=MySQL) [![Github tag](https://badgen.net/github/tag/Naereen/Strapdown.js)](https://github.com/ruanlas/wallet-core-api/tags/)

# wallet-core-api

O objetivo deste projeto é oferecer recursos para a gestão financeira pessoal.

## Índice
<!--ts-->
   * [Tecnologias utilizadas neste projeto](#tecnologias-utilizadas-neste-projeto)
   * [Pré-requisitos](#pré-requisitos)
   * [Ambiente de desenvolvimento](#ambiente-de-desenvolvimento)
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
