ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine as gobuilder
ENV GIN_MODE=release
#RUN apk add --no-cache git
WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./messanger ./cmd/messanger
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./healthy ./cmd/healthy

FROM gcr.io/distroless/static AS final
#FROM alpine:3.16
ENV GIN_MODE=release
#RUN apk --no-cache add ca-certificates tzdata
ENV TZ="Europe/Minsk"
WORKDIR /usr/source/app/
COPY --from=gobuilder /app/messanger /bin/messanger
COPY --from=gobuilder /app/healthy /bin/healthy
