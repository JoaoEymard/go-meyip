# Contrato webservice juno

### [POST] /remessa/:id

#### Request

```
HEADERS {
  "dominio": String,
  "aplicacao": String
}
```

#### Response

```
200 {
  "message": String
}
```

```
404, 500 {
  "error": Object
  "message": String
}
```
