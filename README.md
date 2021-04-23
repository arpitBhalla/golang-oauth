# OAuth2.0 API

- **Golang**
- **MongoDB** for storing data
- **Redis** to Store JWT Metadata

## Base URL

http://34.69.239.180

## Authentication Header

 `Bearer: <Token>`

## Endpoints


### **Register**

> **`POST`**  [**/register**](http://34.69.239.180/register)

**Request Body**

```json
{
  "name": "YOUR NAME",
  "email": "EMAIL",
  "password": "PASSWORD"
}
```

**Response**

```json
{
  "code": 200
}
```


### **Login**

> **`POST`**  [**/login**](http://34.69.239.180/login)

**Request Body**

```json
{
  "email": "EMAIL",
  "password": "PASSWORD"
}
```

**Response**

```json
{
  "code": 200,
  "accessToken": "JWT TOKEN",
  "refreshToken": "JWT TOKEN"
}
```


### **Refresh Token**

> **`POST`**  [**/refresh**](http://34.69.239.180/refresh)

**Request Body**

```json
{
  "refreshToken": "JWT TOKEN"
}
```

**Response**

```json
{
  "accessToken": "JWT TOKEN"
}
```


### **Logout**

> **`POST`**  [**/logout**](http://34.69.239.180/logout)



**Response**

```json
{
  "code": 200
}
```


### **Get your profile**

> **`GET`**  [**/profile**](http://34.69.239.180/profile)



**Response**

```json
{
  "code": 200,
  "_id": "MONGODB ID",
  "name": "NAME",
  "email": "EMAIL"
}
```


### **Dump all profile**

> **`GET`**  [**/all**](http://34.69.239.180/all)



**Response**

```json
{
  "code": 200,
  "data": [
    {
      "_id": "MONGODB ID",
      "name": "NAME",
      "email": "EMAIL"
    }
  ]
}
```

