### DOCS

- GIN: https://github.com/gin-gonic/gin
- GORM: https://gorm.io/docs/index.html
- REST CLIENT: https://github.com/Huachao/vscode-restclient

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
CompileDaemon --build="go build main.go" --command=./main
```

### DOCKER COMPOSE

```shell
docker build -t app .

docker run app

docker compose up -d


docker build -t docker-gs-ping . && docker run -d -t  docker-gs-ping
```

### TEST

```shell
grc go test -v ./tests/...
```
