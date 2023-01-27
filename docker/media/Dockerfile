FROM golang:alpine3.16 as gobuilder

ENV GIN_MODE=release

RUN apk add git
RUN apk add --update gcc musl-dev

WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN go build -o ./media ./cmd/media

FROM alpine:3.16
ENV GIN_MODE=release
RUN apk add --no-cache tzdata
ENV TZ="Europe/Minsk"

########################
# TRANSMISSION #
RUN apk --no-cache --update add transmission-daemon transmission-cli \
  && addgroup transmission users
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

COPY --from=gobuilder /app/media /bin/media
ADD ./wait-for /usr/local/sbin/wait-for
