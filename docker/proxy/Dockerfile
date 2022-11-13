FROM golang:alpine3.16 as gobuilder

ENV GIN_MODE=release

RUN apk add git
RUN apk add --update gcc musl-dev

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN go build -o ./proxy ./cmd/proxy

FROM alpine:3.16
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates tzdata curl
ENV TZ="Europe/Minsk"
EXPOSE 53/udp 53 853
COPY --from=gobuilder /app/proxy /bin/proxy
ADD ./wait-for /usr/local/sbin/wait-for
