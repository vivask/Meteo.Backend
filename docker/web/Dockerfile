FROM golang:alpine3.16 as gobuilder

ENV GIN_MODE=debug

RUN apk add git
RUN apk add --update gcc musl-dev

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN go build -o ./web ./cmd/web

FROM alpine:3.16
ENV GIN_MODE=debug
RUN apk --no-cache add ca-certificates tzdata curl
ENV TZ="Europe/Minsk"
WORKDIR /usr/source/app/
RUN mkdir -p firmware
COPY ./ui ./ui/
COPY --from=gobuilder /app/web /bin/web
ADD ./wait-for /usr/local/sbin/wait-for
