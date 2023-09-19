FROM golang:1.18-alpine

WORKDIR /go/src
COPY ./src .
COPY src/.air.toml /go/src/.air.toml

RUN apk upgrade --update && \
    apk --no-cache add git && \
    go get -u github.com/gin-gonic/gin && \
    go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
