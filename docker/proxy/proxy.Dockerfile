FROM golang:alpine as gobuilder
#FROM golang:arm32v7/alpine as gobuilder
#FROM golang:arm64v8/alpine as gobuilder

ENV GIN_MODE=debug

RUN apk add git
RUN apk add --update gcc musl-dev

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN go build -o ./proxy ./cmd/proxy

FROM alpine
#FROM arm32v7/alpine
#FROM arm64v7/alpine
ENV GIN_MODE=dbug
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY ./config ./config/
EXPOSE 53
COPY --from=gobuilder /app/proxy /bin/proxy
ADD ./wait-for ./wait-for
CMD ./wait-for $WEB_HOST:$WEB_PORT -- echo "postgres is up" && proxy start -m