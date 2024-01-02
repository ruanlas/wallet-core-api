# Wallet Core
API que dispões de recursos para gerenciar as finanças pessoais

## Version: 0.1.0

---
### /v1/gain

#### GET
##### Summary

Obter uma listagem de Receitas

##### Description

Este endpoint permite obter uma listagem de receitas

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| page_size | query | O número de registros retornados pela busca | No | string |
| page | query | A página que será buscada | No | string |
| month | query | O mês que será filtrado a busca | Yes | string |
| year | query | O ano que será filtrado a busca | Yes | string |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [gservice.GainPaginateResponse](#gservicegainpaginateresponse) |

#### POST
##### Summary

Criar uma Receita

##### Description

Este endpoint permite criar uma receita

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| gain | body | Modelo de criação da receita | Yes | [gservice.CreateRequest](#gservicecreaterequest) |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [gservice.GainResponse](#gservicegainresponse) |

### /v1/gain/{id}

#### DELETE
##### Summary

Remove uma Receita

##### Description

Este endpoint permite remover uma receita

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id | path | Id da receita | Yes | string |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [gain.ResponseDefault](#gainresponsedefault) & { **"message"**: string, **"status"**: integer } |

#### GET
##### Summary

Obter uma Receita

##### Description

Este endpoint permite obter uma receita

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| id | path | Id da receita | Yes | string |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [gservice.GainResponse](#gservicegainresponse) |

#### PUT
##### Summary

Editar uma Receita

##### Description

Este endpoint permite editar uma receita

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| gain | body | Modelo de edição da receita | Yes | [gservice.UpdateRequest](#gserviceupdaterequest) |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [gservice.GainResponse](#gservicegainresponse) |

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

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [gpservice.GainProjectionPaginateResponse](#gpservicegainprojectionpaginateresponse) |

#### POST
##### Summary

Criar uma Receita Prevista

##### Description

Este endpoint permite criar uma receita prevista

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| gain_projection | body | Modelo de criação da receita prevista | Yes | [gpservice.CreateRequest](#gpservicecreaterequest) |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 201 | Created | [gpservice.GainProjectionResponse](#gpservicegainprojectionresponse) |

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

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [gpservice.GainProjectionResponse](#gpservicegainprojectionresponse) |

#### PUT
##### Summary

Editar uma Receita Prevista

##### Description

Este endpoint permite editar uma receita prevista

##### Parameters

| Name | Located in | Description | Required | Schema |
| ---- | ---------- | ----------- | -------- | ------ |
| gain_projection | body | Modelo de edição da receita prevista | Yes | [gpservice.UpdateRequest](#gpserviceupdaterequest) |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [gpservice.GainProjectionResponse](#gpservicegainprojectionresponse) |

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
| gain | body | Modelo de criação da receita | Yes | [gpservice.CreateGainRequest](#gpservicecreategainrequest) |
| X-Access-Token | header | Token de autenticação do usuário | Yes | string |

##### Responses

| Code | Description | Schema |
| ---- | ----------- | ------ |
| 200 | OK | [gpservice.GainResponse](#gpservicegainresponse) |

---
### Models

#### gain.ResponseDefault

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| gain.ResponseDefault | object |  |  |

#### gainprojection.ResponseDefault

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| gainprojection.ResponseDefault | object |  |  |

#### gpservice.CategoryResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | string |  | No |
| id | integer |  | No |

#### gpservice.CreateGainRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| pay_in | string |  | No |
| value | number |  | No |

#### gpservice.CreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category_id | integer |  | No |
| description | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| recurrence | integer |  | No |
| value | number |  | No |

#### gpservice.GainProjectionPaginateResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| current_page | integer |  | No |
| page_limit | integer |  | No |
| records | [ [gpservice.GainProjectionResponse](#gpservicegainprojectionresponse) ] |  | No |
| total_pages | integer |  | No |
| total_records | integer |  | No |

#### gpservice.GainProjectionResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [gpservice.CategoryResponse](#gpservicecategoryresponse) |  | No |
| description | string |  | No |
| id | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| recurrence | integer |  | No |
| value | number |  | No |

#### gpservice.GainResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [gpservice.CategoryResponse](#gpservicecategoryresponse) |  | No |
| description | string |  | No |
| gain_projection_id | string |  | No |
| id | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| value | number |  | No |

#### gpservice.UpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category_id | integer |  | No |
| description | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| value | number |  | No |

#### gservice.CategoryResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | string |  | No |
| id | integer |  | No |

#### gservice.CreateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category_id | integer |  | No |
| description | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| value | number |  | No |

#### gservice.GainPaginateResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| current_page | integer |  | No |
| page_limit | integer |  | No |
| records | [ [gservice.GainResponse](#gservicegainresponse) ] |  | No |
| total_pages | integer |  | No |
| total_records | integer |  | No |

#### gservice.GainResponse

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category | [gservice.CategoryResponse](#gservicecategoryresponse) |  | No |
| description | string |  | No |
| gain_projection_id | string |  | No |
| id | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| value | number |  | No |

#### gservice.UpdateRequest

| Name | Type | Description | Required |
| ---- | ---- | ----------- | -------- |
| category_id | integer |  | No |
| description | string |  | No |
| is_passive | boolean |  | No |
| pay_in | string |  | No |
| value | number |  | No |
