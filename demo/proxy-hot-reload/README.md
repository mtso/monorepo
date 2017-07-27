# hot reload proxy

[summary]::
A reverse proxy that flip-flops between routes.

`GET /flip` => triggers the server to reload its routes and only respond 200 to `/flop`,
400 for everything else including `/flip`

`GET /flop` => triggers the server to reload its routes and only respond 200 to `/flip`,
400 for everything else including `/flop`
