# OAuth2.0 API

- **Golang**
- **MongoDB** for storing data
- **Redis** to Store JWT Metadata

## Base URL

https://nitkkr-online.el.r.appspot.com/

## Authentication Header

 `Bearer: <Token>`

## Endpoints


### **Register**

> **`POST`**  [**/register**](https://nitkkr-online.el.r.appspot.com//register)

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

> **`POST`**  [**/login**](https://nitkkr-online.el.r.appspot.com//login)

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

> **`POST`**  [**/refresh**](https://nitkkr-online.el.r.appspot.com//refresh)

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

> **`POST`**  [**/logout**](https://nitkkr-online.el.r.appspot.com//logout)



**Response**

```json
{
  "code": 200
}
```


### **Get your profile**

> **`GET`**  [**/profile**](https://nitkkr-online.el.r.appspot.com//profile)



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

> **`GET`**  [**/all**](https://nitkkr-online.el.r.appspot.com//all)



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

