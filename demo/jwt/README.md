# jwt

[summary]::
Demo jwt signing and verification using `github.com/dgrijalva/jwt-go`.

## test plan
1. torus run ./sign -payload '{"message":"hello, world"}'
    > asdfsa.fdsafdsa.fdsafdsa
2. torus run ./verify -jwt 'asdfsa.fdsafdsa.fdsafdsa'
    > valid/invalid
    > payload: {"message":"hello, world"}
