ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine as gobuilder
ENV GIN_MODE=release
WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./radio ./cmd/radio
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./healthy ./cmd/healthy

FROM alpine:3.16
ENV GIN_MODE=release
RUN apk add --no-cache tzdata
ENV TZ="Europe/Minsk"
# dump1090 install
RUN apk add --no-cache rtl-sdr ncurses
# install build packages
RUN \
apk add --no-cache --virtual=build-dependencies \
  git \
  curl \
  g++ \
  gcc \
  librtlsdr-dev \
  ncurses-dev \
  cmake \
  make && \
cd /tmp && \
#git clone https://github.com/antirez/dump1090 && \
git clone https://github.com/flightaware/dump1090.git && \
cd dump1090 && \
make BLADERF=no HACKRF=no LIMESDR=no && \
cp dump1090 /bin && \
# cleanup
apk del --purge \
  build-dependencies && \
rm -rf \
  /root/* \
  /tmp/* \
  /usr/lib/*.la

EXPOSE 30001/tcp 30002/tcp 30003/tcp 30004/tcp 30005/tcp 30104/tcp

COPY --from=gobuilder /app/radio /bin/radio
COPY --from=gobuilder /app/healthy /bin/healthy
