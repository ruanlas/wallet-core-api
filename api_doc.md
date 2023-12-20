# Wallet Core
API que dispões de recursos para gerenciar as finanças pessoais

## Version: 0.1.0

---
### /v1/gain-projection

#### GET
##### Summary

Obter uma listagem de Receitas Previstas

##### Description

Este endpoint permite obter uma listagem de receitas previstas

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| page_size | query | O número de registros retornados pela busca | No | string |
| page | query | A página que será buscada | No | string |
| month | query | O mês que será filtrado a busca | Yes | string |
| year | query | O ano que será filtrado a busca | Yes | string |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |
| X-Userinfo | header | Informações do usuário em base64 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [service.GainProjectionPaginateResponse](#servicegainprojectionpaginateresponse) |

#### POST
##### Summary

Criar uma Receita Prevista

##### Description

Este endpoint permite criar uma receita prevista

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| gain_projection | body | Modelo de criação da receita prevista | Yes | [service.CreateRequest](#servicecreaterequest) |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |
| X-Userinfo | header | Informações do usuário em base64 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [service.GainProjectionResponse](#servicegainprojectionresponse) |

### /v1/gain-projection/{id}

#### DELETE
##### Summary

Remove uma Receita Prevista

##### Description

Este endpoint permite remover uma receita prevista

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id | path | Id da receita prevista | Yes | string |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |
| X-Userinfo | header | Informações do usuário em base64 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [gainprojection.ResponseDefault](#gainprojectionresponsedefault) & { **"message"**: string, **"status"**: integer } |

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

#### PUT
##### Summary

Editar uma Receita Prevista

##### Description

Este endpoint permite editar uma receita prevista

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| gain_projection | body | Modelo de edição da receita prevista | Yes | [service.UpdateRequest](#serviceupdaterequest) |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |
| X-Userinfo | header | Informações do usuário em base64 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [service.GainProjectionResponse](#servicegainprojectionresponse) |

### /v1/gain-projection/{id}/create-gain

#### POST
##### Summary

Realizar uma Receita Prevista

##### Description

Este endpoint permite realizar uma receita que foi prevista

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id | path | Id da receita prevista | Yes | string |
| gain | body | Modelo de criação da receita | Yes | [service.CreateGainRequest](#servicecreategainrequest) |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |
| X-Userinfo | header | Informações do usuário em base64 | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [service.GainResponse](#servicegainresponse) |

---
### Models

#### gainprojection.ResponseDefault

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| gainprojection.ResponseDefault | object |  |  |

#### service.CategoryResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | string |  | No |
| id | integer |  | No |

#### service.CreateGainRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| pay_in | string |  | No |
| value | number |  | No |

#### service.CreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category_id | integer |  | No |
| description | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| recurrence | integer |  | No |
| value | number |  | No |

#### service.GainProjectionPaginateResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| current_page | integer |  | No |
| page_limit | integer |  | No |
| records | [ [service.GainProjectionResponse](#servicegainprojectionresponse) ] |  | No |
| total_pages | integer |  | No |
| total_records | integer |  | No |

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

#### service.GainResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [service.CategoryResponse](#servicecategoryresponse) |  | No |
| description | string |  | No |
| gain_projection_id | string |  | No |
| id | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| value | number |  | No |

#### service.UpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category_id | integer |  | No |
| description | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| value | number |  | No |
