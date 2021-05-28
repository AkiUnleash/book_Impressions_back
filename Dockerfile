# Herokuデプロイ用
FROM golang:1.15.2-alpine

RUN apk update && apk add git

RUN mkdir /go/src/api
WORKDIR /go/src/api

ADD ./api /go/src/api

RUN GOOS=linux GOARCH=amd64 go build -o /main

FROM alpine:3.9

COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main"]

# RUN ./api/main