FROM golang:alpine:3.16.2 as gobuilder
#FROM golang:arm32v7/alpine:3.13 as gobuilder
#FROM golang:arm64v8/alpine:3.13 as gobuilder

ENV GIN_MODE=debug

RUN apk add git
RUN apk add --update gcc musl-dev

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN go build -o ./proxy ./cmd/proxy

FROM alpine:3.16.2
#FROM arm32v7/alpine:3.13
#FROM arm64v7/alpine:3.13
ENV GIN_MODE=dbug
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY ./config ./config/
EXPOSE 53
COPY --from=gobuilder /app/proxy /bin/proxy
ADD ./wait-for ./wait-for
ENV DATABASE_TYPE="postgres"
ENV DATABASE_URL="host=db port=5432 user=admin dbname=meteo password=P@55w0rd sslmode=disable"
CMD ./wait-for $DATABASE_HOST:$DATABASE_PORT -- echo "postgres is up" && proxy migrate && proxy start