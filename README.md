# Gawds Auth Golang

## API

### Base URL

> https://nitkkr-online.el.r.appspot.com/

### Endpoints

#### Register

> **`POST` /register**

<img src="img/register.png">

#### Login

> **`POST` /login**

<img src="img/login.png">

#### Get Current User

> **`GET` /profile**

<img src="img/profile.png">

#### Logout

> **`POST` /logout**

<img src="img/logout.png">

#### Get All Users

> **`GET` /all**

<img src="img/all.png">

**Response**

```json
{
  "code": 400,
  "result": [
    {
      "_id": "60803df858302ce4992142bb",
      "name": "Arpit Bhalla",
      "email": "arpitbhalla2001@gmail.com",
      "password": "demo1234"
    }
  ]
}
```
