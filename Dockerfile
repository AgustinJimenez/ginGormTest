FROM golang:latest
RUN go install github.com/githubnemo/CompileDaemon@latest
WORKDIR /app
COPY . .
RUN go mod download -x
ENTRYPOINT CompileDaemon --build="go build /app/main.go" --command="/app/main"