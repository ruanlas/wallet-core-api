# Wallet Core
API que dispões de recursos para gerenciar as finanças pessoais

## Version: 0.1.0

---
### /v1/gain-projection

#### POST
##### Summary

Criar uma Receita Prevista

##### Description

Este endpoint permite criar uma receita prevista

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| gain_projection | body | Modelo de criação da receita | Yes | [service.CreateRequest](#servicecreaterequest) |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |
| X-Userinfo | header | Informações do usuário em base64 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [service.GainProjectionResponse](#servicegainprojectionresponse) |

### /v1/gain-projection/{id}

#### GET
##### Summary

Obter uma Receita Prevista

##### Description

Este endpoint permite obter uma receita prevista

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id | path | Id da receita prevista | Yes | string |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |
| X-Userinfo | header | Informações do usuário em base64 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [service.GainProjectionResponse](#servicegainprojectionresponse) |

---
### Models

#### service.CategoryResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | string |  | No |
| id | integer |  | No |

#### service.CreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category_id | integer |  | No |
| description | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| recurrence | integer |  | No |
| value | number |  | No |

#### service.GainProjectionResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [service.CategoryResponse](#servicecategoryresponse) |  | No |
| description | string |  | No |
| id | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| recurrence | integer |  | No |
| value | number |  | No |
