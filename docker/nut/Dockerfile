ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine as gobuilder
ENV GIN_MODE=release
#RUN apk add --no-cache git
WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./nut ./cmd/nut
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./upssched-cmd ./cmd/upssched-cmd
#RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./healthy ./cmd/healthy

FROM alpine:3.16
ENV GIN_MODE=release
RUN apk add --no-cache tzdata
ENV TZ="Europe/Minsk"

########################
# NUT #
ARG NUT_VERSION=2.8.0-r4
ENV USER=nut \
    GROUP=nut
RUN echo '@edge http://dl-cdn.alpinelinux.org/alpine/edge/community' \
    >>/etc/apk/repositories && \
  apk add --update nut@edge=$NUT_VERSION \
    libcrypto1.1 libssl1.1 musl net-snmp-libs
RUN chgrp $GROUP /etc/nut/*
RUN chmod 640 /etc/nut/*
RUN mkdir -p -m 2750 /dev/shm/nut \
  && chown -R $USER:$GROUP /dev/shm/nut \
  && mkdir -p /var/run/nut \
  && echo 0 > /var/run/nut/upsd.pid \
  && chown -R $USER:$GROUP /var/run/nut \
  && echo 0 > /var/run/upsmon.pid
RUN sed -i "s/none/standalone/g" /etc/nut/nut.conf
RUN echo 'LISTEN 127.0.0.1 3493' >> /etc/nut/upsd.conf
RUN echo '' >> /etc/nut/ups.conf \
  #&& echo 'pollonly = "enabled"' >> /etc/nut/ups.conf \
  #&& echo '' >> /etc/nut/ups.conf \
  && echo '[powercom]' >> /etc/nut/ups.conf \
  && echo 'driver = usbhid-ups' >> /etc/nut/ups.conf \
  && echo 'port = auto' >> /etc/nut/ups.conf \
  && echo 'desc = "WOW-700U"' >> /etc/nut/ups.conf
  #&& echo 'vendorid = 0d9f' >> /etc/nut/ups.conf \
  #&& echo 'productid = 0004' >> /etc/nut/ups.conf
  #&& echo 'pollinterval = 10' >> /etc/nut/ups.conf
#RUN echo 'SUBSYSTEM=="usb", ATTR{idVendor}=="0d9f", MODE="0666"' >> /etc/udev/rules.d/50-ups-powercom.rules \
#  && echo 'SUBSYSTEM=="usb", ACTION=="add", ATTR{idVendor}=="0d9f", ATTR{idProduct}=="0004", RUN+="/usr/sbin/upsdrvctl -u root shutdown"' >> /etc/udev/rules.d/50-ups-powercom.rules
RUN --mount=type=secret,id=nut_env source /run/secrets/nut_env \
  && echo '['$API_USER']' >> /etc/nut/upsd.users \
  && echo 'password = '$API_PASSWORD >> /etc/nut/upsd.users \
  && echo 'upsmon = master' >> /etc/nut/upsd.users \
  && echo 'actions = SET' >> /etc/nut/upsd.users \
  && echo 'actions = FSD' >> /etc/nut/upsd.users \
  && echo 'instcmds = ALL' >> /etc/nut/upsd.users \
  && echo 'NOTIFYCMD /usr/sbin/upssched' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYMSG ONLINE  "UPS %s on line power"' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYMSG ONBATT  "UPS %s on battery"' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYMSG LOWBATT  "UPS %s battery is low"' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYMSG FSD    "UPS %s: forced shutdown in progress"' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYMSG COMMOK  "Communications with UPS %s established"' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYMSG COMMBAD  "Communications with UPS %s lost"' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYMSG SHUTDOWN  "Auto logout and shutdown proceeding"' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYMSG REPLBATT  "UPS %s battery needs to be replaced"' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYMSG NOCOMM  "UPS %s is unavailable"' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYMSG NOPARENT  "upsmon parent process died - shutdown impossible"' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYFLAG ONLINE EXEC' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYFLAG ONBATT EXEC' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYFLAG LOWBATT EXEC' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYFLAG FSD EXEC' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYFLAG COMMOK EXEC' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYFLAG COMMBAD EXEC' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYFLAG SHUTDOWN EXEC' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYFLAG REPLBATT EXEC' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYFLAG NOCOMM EXEC' >> /etc/nut/upsmon.conf \
  && echo 'NOTIFYFLAG NOPARENT EXEC' >> /etc/nut/upsmon.conf \
  && echo 'MONITOR powercom@localhost 1' $API_USER $API_PASSWORD 'master' >> /etc/nut/upsmon.conf \
  && echo 'RUN_AS_USER' $USER >> /etc/nut/upsmon.conf
RUN echo 'PIPEFN /tmp/upspipe' >> /etc/nut/upssched.conf \
  && echo 'LOCKFN /tmp/upslock' >> /etc/nut/upssched.conf \
  && echo 'AT COMMBAD * EXECUTE commbad' >> /etc/nut/upssched.conf \
  && echo 'AT COMMOK * EXECUTE commok' >> /etc/nut/upssched.conf \
  && echo 'AT NOCOMM * EXECUTE nocomm' >> /etc/nut/upssched.conf \
  && echo 'AT ONBATT * EXECUTE powerout' >> /etc/nut/upssched.conf \
  && echo 'AT ONBATT * START-TIMER shutdownnow 1200' >> /etc/nut/upssched.conf \
  && echo 'AT LOWBATT * EXECUTE shutdowncritical' >> /etc/nut/upssched.conf \
  && echo 'AT ONLINE * CANCEL-TIMER shutdownnow' >> /etc/nut/upssched.conf \
  && echo 'AT ONLINE * EXECUTE powerup' >> /etc/nut/upssched.conf \
  && echo 'AT REPLBATT * EXECUTE replbatt' >> /etc/nut/upssched.conf \
  && echo 'AT NOPARENT * EXECUTE noparent' >> /etc/nut/upssched.conf

COPY --from=gobuilder /app/nut /bin/nut
COPY --from=gobuilder /app/upssched-cmd /usr/bin/upssched-cmd
#COPY --from=gobuilder /app/healthy /bin/healthy
