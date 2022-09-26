FROM golang:alpine as gobuilder
#FROM golang:arm32v7/alpine as gobuilder
#FROM golang:arm64v8/alpine as gobuilder

ENV GIN_MODE=debug

RUN apk add git
RUN apk add --update gcc musl-dev bash mc

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN go build -o ./web ./cmd/web

FROM alpine
#FROM arm32v7/alpine
ENV GIN_MODE=dbug
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY ./config ./config/

COPY --from=gobuilder /app/web /bin/web
ADD ./wait-for ./wait-for
ENV DATABASE_TYPE="postgres"
ENV DATABASE_URL="host=db port=5432 user=admin dbname=meteo password=P@55w0rd sslmode=disable"
CMD ./wait-for $DATABASE_HOST:$DATABASE_PORT -- echo "postgres is up" && xu4 start