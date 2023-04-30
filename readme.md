### DOCS

- GIN: https://github.com/gin-gonic/gin
- GORM: https://gorm.io/docs/index.html

### generate .env file from .env.example

```
cp .env.example .env
```

### RUN MIGRATIONS

```shell
go run migrations/migrate.go
```

### RUN SERVER

```shell
CompileDaemon -command="./go_practice"
```

### TEST

```shell
go test ./tests/...
```
