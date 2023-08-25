ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine as gobuilder
ENV GIN_MODE=release
#RUN apk add --no-cache git
WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./cluster ./cmd/cluster
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./healthy ./cmd/healthy

FROM gcr.io/distroless/static AS final
ENV GIN_MODE=release
#RUN apk --no-cache add ca-certificates tzdata
ENV TZ="Europe/Minsk"
COPY --from=gobuilder /app/cluster /bin/cluster
COPY --from=gobuilder /app/healthy /bin/healthy
