To start project use command
```
docker compose up -d
```
and run migration
```
goose -dir migrations postgres "user=postgres password=qwerty dbname=blog host=127.0.0.1 port=15432 sslmode=disable" up
```

Endpoints:
```
Auth:
Register user: /register POST
Login: /login POST

Post:
Create post: /posts POST
Get author post list: /posts/{authorId} GET
```

