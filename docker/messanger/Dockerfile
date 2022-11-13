FROM golang:alpine3.16 as gobuilder

ENV GIN_MODE=release

RUN apk add git
RUN apk add --update gcc musl-dev

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN go build -o ./messanger ./cmd/messanger

FROM alpine:3.16
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates tzdata curl
ENV TZ="Europe/Minsk"
COPY --from=gobuilder /app/messanger /bin/messanger
#ADD ./wait-for /usr/local/sbin/wait-for
