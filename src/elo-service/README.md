# ELO leaderboard service

[summary]::
League management service that calculates player ELO ratings.

## User Stories
- I can create a new league.
- I can add a new game with a winner and loser to a league.
- I can view a list of the players in a league and their ELO scores.
- I can view a list of all the games played in a league.

## Routes
```
GET  /           -- serve client app
GET  /docs       -- get the API docs
GET  /:league_id -- show league page

GET  /api/info   -- data about the microservice
POST /api/new
POST /api/:league_id
GET  /api/:league_id/players
GET  /api/:league_id/games
```

## League data
```json
{
  "title": "The League Title",
  "id": "119fu89fhi7euafhdjskl"
}
```

## Players (sorted by elo descending)
```json
[
  {
    "username": "wiggs",
    "elo": 345
  },
  ...
]
```

## Games (sorted in reverse chronological order)
```json
[
  {
    "winner": {
      "username": "wiggs",
      "elo": 345
    },
    "loser": {
      "username": "wiggs",
      "elo": 345
    },
    "timestamp": "123T123Z"
  }
  ...
]
```
