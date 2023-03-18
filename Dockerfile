FROM golang:1.15-alpine3.12 AS builder

COPY . /github.com/dmytrodemianchuk/cz-api/
WORKDIR /github.com/dmytrodemianchuk/cz-api/

RUN go mod download
RUN go build -o ./bin/bot cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/dmytrodemianchuk/cz-api/bin/bot .

EXPOSE 80

CMD ["./bot"]