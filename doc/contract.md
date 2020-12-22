# Contrato webservice juno

### 1. [POST] /acao/enviar-remessa

Criar uma ação de worker (trabalhador) para o envio de boletos de uma remessa para a juno.

#### - **_Request_**

```
HEADERS = {
  "x-domain": String,
  "x-app-key": String,
  "x-aplication": String,
  "x-version": String, // 1.0.0
}
```

```
BODY = {
  "id_remessa": Integer,
  "_mid": String,
  "_v": String,
}
```

#### - **_Response_**:

```
STATUS = 200 => {
  "aid": String,
  "_mid": String,
}
```

```
STATUS = 404, 500 => {
  "error": Object
  "message": String
}
```
