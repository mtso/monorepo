# dotnet toy problem

[summary]::
Build a web API server using dotnet and persist data using SQL Server.

## API Specification

### Get a random word

URL Endpoint: `GET /api/word`

Example Request

```
GET /api/word
```

Example Response

```json
{
  "is_success": true,
  "message": null,
  "content": {
    "id": 123,
    "value": "attention",
    "level": 1,
    "created_on": "2016-01-21T13:46:57.48"
  }
}
```

### Get word list

URL Endpoint: `GET /api/words`, parameters: `?level=[integer]`

Example Request (Get all words)

```
GET /api/words
```

Example Response

```json
{
  "is_success": true,
  "message": null,
  "content": [
    {
      "id": 123,
      "value": "attention",
      "level": 1,
      "created_on": "2016-01-21T13:46:57.48"
    },
    {
      "id": 124,
      "value": "magic",
      "level": 1,
      "created_on": "2016-01-21T13:46:57.48"
    }
  ]
}
```

Example Request (Get all level 1 words)

```
GET /api/words?level=1
```

Example Response

```json
{
  "is_success": true,
  "message": null,
  "content": [
    {
      "id": 123,
      "value": "attention",
      "level": 1,
      "created_on": "2016-01-21T13:46:57.48"
    },
    {
      "id": 124,
      "value": "magic",
      "level": 1,
      "created_on": "2016-01-21T13:46:57.48"
    }
  ]
}
```

### Add new word

URL Endpoint: `POST /api/word`

Example Request

```
POST /api/word
Body {
  "value": "fantasy",
  "level": 1
}
```

Example Response

```json
{
  "is_success": true,
  "message": null,
  "content": {
    "id": 125,
    "value": "fantasy",
    "level": 1,
    "created_on": "2017-08-11T13:46:57.48"
  }
}
```

### Update a word

URL Endpoint: `PUT /api/word/:id`

Example Request

```
PUT /api/word/123
Body {
  "level": 2
}
```

Example Response

```json
{
  "is_success": true,
  "message": null,
  "content": {
    "id": 123,
    "value": "attention",
    "level": 2,
    "created_on": "2017-08-11T13:46:57.48"
  }
}
```

### Get specific word

URL Endpoint: `GET /api/word/:id`

Example Request

```
GET /api/word/123
```

Example Response

```json
{
  "is_success": true,
  "message": null,
  "content": {
    "id": 123,
    "value": "attention",
    "level": 2,
    "created_on": "2017-08-11T13:46:57.48"
  }
}
```

### Delete word

URL Endpoint: `DELETE /api/word/:id`

Example Request

```
DELETE /api/word/123
```

Example Response

```json
{
  "is_success": true,
  "message": null,
  "content": {
    "id": 123,
    "value": "attention",
    "level": 2,
    "created_on": "2017-08-11T13:46:57.48"
  }
}
```

## Sample Schema

### Word
name          | datatype  | options
------------- | --------- | ------------------------------ 
id            | bigint    | not null, unique, incrementing
value         | varchar   | not null, unique
level         | int       | not null, indexed
created_on    | Datetime2 | not null

### SQL Sketch

CREATE TABLE IF NOT EXISTS Word (
    id BIGINT NOT NULL UNIQUE IDENTITY(1,1),
    value VARCHAR NOT NULL UNIQUE,
    level INT NOT NULL,

    CONSTRAINT PK_Word PRIMARY KEY CLUSTERED (id),
    CONSTRAINT 
)
