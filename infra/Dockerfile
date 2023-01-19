FROM golang:1.19.5-alpine3.17

RUN apk update && apk add \
    git

RUN go install github.com/cosmtrek/air@v1.40.4

WORKDIR /app

CMD ["air", "-c", ".air.toml"]