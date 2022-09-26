FROM golang:alpine as gobuilder
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
RUN go build -o ./web ./cmd/web

FROM alpine
#FROM arm32v7/alpine
#FROM arm64v7/alpine
ENV GIN_MODE=release
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY ./config ./config/
EXPOSE 80
COPY --from=gobuilder /app/web /bin/web
ADD ./wait-for ./wait-for
ENV DATABASE_TYPE="postgres"
ENV DATABASE_URL="host=db port=5432 user=admin dbname=meteo password=P@55w0rd sslmode=disable"
CMD ./wait-for $DATABASE_HOST:$DATABASE_PORT -- echo "postgres is up" && server migrate && server start