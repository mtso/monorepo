# API

The API is based on a JSON request and response model.
`GET` requests do not expect parameters besides the league ID in the URL.
`POST` requests should JSON-encode the parameters and be sent in the request body.

## Create a League
```url
GET /api/new
```

#### Example Response
```json
{
  "ok": true,
  "league": {
    "id": "7db934381579dee00af58b3f",
    "title": "tea party",
    "createdAt": "2017-07-08T10:34:40.699839Z"
  }
}
```

## Add a Game
```url
POST /api/:id/
```

#### Example Request
```json
{
  "winner": "username1",
  "loser": "username2"
}
```

#### Example Response
```
{
  "ok": true,
  "game": {
    "createdAt": "2017-07-08T16:15:59.560923Z",
    "winner": {
      "username": "username1",
      "elo": 1519
    },
    "loser": {
      "username": "username2",
      "elo": 1324
    }
  }
}
```

## Get League Info
```url
GET /api/:id
```

#### Example Response
```json
{
  "ok": true,
  "league": {
    "id": "7db934381579dee00af58b3f",
    "title": "tea party",
    "createdAt": "2017-07-08T10:34:40.699839Z"
  }
}
```

## Get Game History of a League
```url
GET /api/:id/games
```

Games are sorted by reverse chronological order.

#### Example Response
```json
{
  "ok": true,
  "games": [
    {
      "createdAt": "2017-07-08T16:15:59.560923Z",
      "winner": {
        "username": "username1",
        "elo": 1519
      },
      "loser": {
        "username": "username2",
        "elo": 1324
      }
    },
    ...
  ]
}
```

## Get the Players in a League
```url
GET /api/:id/games
```

Players are sorted by ELO rating in descending order.

#### Example Response
```json
{
  "ok": true,
  "players": [
    {
      "username": "username1",
      "elo": 1519
    },
    {
      "username": "username2",
      "elo": 1324
    },
    ...
  ]
}
```
