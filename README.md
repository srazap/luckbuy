# luckbuy
LuckBuy Home Assignment

## How to run

1) Install postgres database
2) Run `go run main.go` to run API server


### Example

```
1) Signup API


curl --location --request POST 'http://localhost:8080/api/v1/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "nithya@gmail.com",
    "password": "nithiya123",
    "refferral_code": "zahid@gmail.com"
}'
```

```
2) Login

curl --location --request POST 'http://localhost:8080/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "gopi@gmail.com",
    "password": "gopi123"
}'

```

```
3) MyPoints

curl --location --request GET 'http://localhost:8080/api/v1/points' \
--header 'session_id: Z29waUBnbWFpbC5jb20xNjQ0ODQ5OTE3'
```

```
4) Leaderboard

curl --location --request GET 'http://localhost:8080/api/v1/leaderboard' \
--header 'session_id: Z29waUBnbWFpbC5jb20xNjQ0ODQ5NjE5'
```

```
5) Logout

curl --location --request GET 'http://localhost:8080/api/v1/logout' \
--header 'session_id: Z29waUBnbWFpbC5jb20xNjQ0ODQ5NjE5'
```