# Gawds Auth Golang

## API

> `Endpoint` https://

### Get all users

> **`GET` /user**

```json
{
  "code": "200",
  "results": [{}]
}
```

### Get user by ID

> **`GET` /user/{id}**

```json
{
  "code": "200",
  "result": {}
}
```

### New User

> **`POST` /user**

**Body:**

```json
{}
```

**Response:**

```json
{
  "code": "200",
  "message": "OK"
}
```

### Update User by ID

> **`PUT` /user/{id}**

**Body:**

```json
{}
```

**Response:**

```json
{}
```

### Delete User by ID

> **`DELETE` /user/{id}**

```json
{
  "code": "200",
  "message": "OK"
}
```
