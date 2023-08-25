ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine as gobuilder
ENV GIN_MODE=release
#RUN apk add --no-cache git
WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./media ./cmd/media
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./healthy ./cmd/healthy
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./flactool ./cmd/flactool

FROM alpine:3.16
ENV GIN_MODE=release
RUN apk add --no-cache tzdata
ENV TZ="Europe/Minsk"

########################
# TRANSMISSION #
RUN apk --no-cache --update add transmission-daemon transmission-cli
RUN addgroup transmission users
ADD ./docker/media/settings.json /etc/transmission-daemon/settings.json
RUN echo 'net.core.rmem_max=4194304' >> /etc/sysctl.conf \
  && echo 'net.core.wmem_max=1048576' >> /etc/sysctl.conf
EXPOSE 9091/tcp 51413

########################
# SAMBA #
RUN apk --no-cache --update add samba
RUN adduser -D -G users -H -S -g 'Samba User' -h /tmp smbuser \
  && mkdir -p /media/samba
RUN --mount=type=secret,id=samba_env source /run/secrets/samba_env && echo -ne "${SMB_PASS}\n${SMB_PASS}\n" | smbpasswd -a -s smbuser
COPY ./docker/media/smb.conf /etc/samba/smb.conf
EXPOSE 137/udp 138/udp 139 445

########################
# FLACTOOL #
# install build packages
RUN \
apk add --no-cache --virtual=build-dependencies \
  git \
  curl \
  g++ \
  gcc \
  libxml2-dev \
  cmake \
  make && \
# install runtime packages
apk --no-cache --update add \
  libxml2 \
  mc \
  ffmpeg \
  sox \
  flac && \
# compile shntool
mkdir -p \
  /tmp/shntool && \
curl -o \
  /tmp/shntool-src-tar.gz -L \
  http://shnutils.freeshell.org/shntool/dist/src/shntool-3.0.10.tar.gz && \
tar xf /tmp/shntool-src-tar.gz -C \
  /tmp/shntool --strip-components=1 && \
cd /tmp/shntool && \
./configure \
  --infodir=/usr/share/info \
  --localstatedir=/var \
  --mandir=/usr/share/man \
  --prefix=/usr \
  --sysconfdir=/etc && \
make && \
make install && \
# compile sacd_extract
cd /root &&\
git clone https://github.com/sacd-ripper/sacd-ripper.git && \
cd sacd-ripper/tools/sacd_extract/ && \
cmake . && \
make && \
cp ./sacd_extract /bin/sacd_extract && \
# cleanup
apk del --purge \
  build-dependencies && \
rm -rf \
  /root/* \
  /tmp/* \
  /usr/lib/*.la

COPY --from=gobuilder /app/media /bin/media
COPY --from=gobuilder /app/healthy /bin/healthy
COPY --from=gobuilder /app/flactool /bin/flactool
