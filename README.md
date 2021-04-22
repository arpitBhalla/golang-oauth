# README

## API

### Base URL

> https://nitkkr-online.el.r.appspot.com/

### Authentication Header

> `Bearer: <Token>`

### Endpoints

#### Register

> `POST`  **/register**

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


#### Login

> `POST`  **/login**

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


#### Refresh Token

> `POST`  **/refresh**

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


#### Logout

> `POST`  **/logout**

null

**Response**

```json
{
  "code": 200
}
```


#### Get your profile

> `GET`  **/profile**

null

**Response**

```json
{
  "code": 200,
  "_id": "MONGODB ID",
  "name": "NAME",
  "email": "EMAIL"
}
```


#### Get your profile

> `GET`  **/all**

null

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

