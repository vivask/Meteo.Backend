ARG GO_VERSION=1.19

FROM golang:${GO_VERSION}-alpine as gobuilder
ENV GIN_MODE=release
#RUN apk add --no-cache git
WORKDIR /app

ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum
RUN go mod download

ADD . /app
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./radius ./cmd/radius
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o ./healthy ./cmd/healthy

FROM alpine:3.16
ENV GIN_MODE=release
RUN apk add --no-cache tzdata
ENV TZ="Europe/Minsk"

########################
# FREERADIUS #
RUN apk --no-cache --update add freeradius freeradius-postgresql freeradius-eap freeradius-utils openssl curl
ADD ./docker/radius/etc/raddb/certs/ca.cnf /etc/raddb/certs/ca.cnf
ADD ./docker/radius/etc/raddb/certs/server.cnf /etc/raddb/certs/server.cnf
ADD ./docker/radius/etc/raddb/certs/client.cnf /etc/raddb/certs/client.cnf
ADD ./docker/radius/etc/raddb/certs/inner-server.cnf /etc/raddb/certs/inner-server.cnf
ADD ./docker/radius/etc/raddb/certs/crl.cnf /etc/raddb/certs/crl.cnf
ADD ./docker/radius/etc/raddb/clients.conf /etc/raddb/clients.conf
ADD ./docker/radius/etc/raddb/mods-available/eap /etc/raddb/mods-available/eap
ADD ./docker/radius/etc/raddb/mods-available/sql /etc/raddb/mods-available/sql
ADD ./docker/radius/etc/raddb/sites-available/default /etc/raddb/sites-available/default
ADD ./docker/radius/etc/raddb/sites-available/inner-tunnel /etc/raddb/sites-available/inner-tunnel
ADD ./docker/radius/etc/raddb/sites-available/tls /etc/raddb/sites-available/tls
ADD ./docker/secrets/n2.radius.env /tmp/radius.env
RUN export $(cat /tmp/radius.env | grep -v '^#' | xargs) \
#CA certificate config
  && sed -i "s/whatever/${CERT_PASS}/g" /etc/raddb/certs/ca.cnf \
  && sed -i "s/<country>/${CERT_COUNTRY_CODE}/g" /etc/raddb/certs/ca.cnf \
  && sed -i "s/<state>/${CERT_STATE}/g" /etc/raddb/certs/ca.cnf \
  && sed -i "s/<city>/${CERT_LOCALITY}/g" /etc/raddb/certs/ca.cnf \
  && sed -i "s/<office>/${CERT_ORGANIZATION}/g" /etc/raddb/certs/ca.cnf \
  && sed -i "s/<email>/${CERT_EMAIL}/g" /etc/raddb/certs/ca.cnf \
  && sed -i "s/<common>/${CERT_CN}/g" /etc/raddb/certs/ca.cnf \
#Server certificate config
  && sed -i "s/whatever/${CERT_PASS}/g" /etc/raddb/certs/server.cnf \
  && sed -i "s/<country>/${CERT_COUNTRY_CODE}/g" /etc/raddb/certs/server.cnf \
  && sed -i "s/<state>/${CERT_STATE}/g" /etc/raddb/certs/server.cnf \
  && sed -i "s/<city>/${CERT_LOCALITY}/g" /etc/raddb/certs/server.cnf \
  && sed -i "s/<office>/${CERT_ORGANIZATION}/g" /etc/raddb/certs/server.cnf \
  && sed -i "s/<email>/${CERT_EMAIL}/g" /etc/raddb/certs/server.cnf \
  && sed -i "s/<common>/${CERT_CN}/g" /etc/raddb/certs/server.cnf \
  && sed -i "s/<ns1>/${CERT_DNS1}/g" /etc/raddb/certs/server.cnf \
  && sed -i "s/<ns2>/${CERT_DNS2}/g" /etc/raddb/certs/server.cnf \
#Client certificate config
  && sed -i "s/whatever/${CERT_PASS}/g" /etc/raddb/certs/client.cnf \
  && sed -i "s/<country>/${CERT_COUNTRY_CODE}/g" /etc/raddb/certs/client.cnf \
  && sed -i "s/<state>/${CERT_STATE}/g" /etc/raddb/certs/client.cnf \
  && sed -i "s/<city>/${CERT_LOCALITY}/g" /etc/raddb/certs/client.cnf \
  && sed -i "s/<office>/${CERT_ORGANIZATION}/g" /etc/raddb/certs/client.cnf \
  && sed -i "s/<email>/${CERT_EMAIL}/g" /etc/raddb/certs/client.cnf \
  && sed -i "s/<client_cn>/${CERT_CLIENT_CN}/g" /etc/raddb/certs/client.cnf \
#Inner server certificate config
  && sed -i "s/whatever/${CERT_PASS}/g" /etc/raddb/certs/inner-server.cnf \
  && sed -i "s/<country>/${CERT_COUNTRY_CODE}/g" /etc/raddb/certs/inner-server.cnf \
  && sed -i "s/<state>/${CERT_STATE}/g" /etc/raddb/certs/inner-server.cnf \
  && sed -i "s/<city>/${CERT_LOCALITY}/g" /etc/raddb/certs/inner-server.cnf \
  && sed -i "s/<office>/${CERT_ORGANIZATION}/g" /etc/raddb/certs/inner-server.cnf \
  && sed -i "s/<email>/${CERT_EMAIL}/g" /etc/raddb/certs/inner-server.cnf \
  && sed -i "s/<common>/${CERT_CN}/g" /etc/raddb/certs/inner-server.cnf \
#CRL certificate config
  && sed -i "s/whatever/${CERT_PASS}/g" /etc/raddb/certs/crl.cnf \
  && sed -i "s/<country>/${CERT_COUNTRY_CODE}/g" /etc/raddb/certs/crl.cnf \
  && sed -i "s/<state>/${CERT_STATE}/g" /etc/raddb/certs/crl.cnf \
  && sed -i "s/<city>/${CERT_LOCALITY}/g" /etc/raddb/certs/crl.cnf \
  && sed -i "s/<office>/${CERT_ORGANIZATION}/g" /etc/raddb/certs/crl.cnf \
  && sed -i "s/<email>/${CERT_EMAIL}/g" /etc/raddb/certs/crl.cnf \
  && sed -i "s/<common>/${CERT_CN}/g" /etc/raddb/certs/crl.cnf \
#TLS config
  && sed -i "s/whatever/${CERT_PASS}/g" /etc/raddb/sites-available/tls \
  && sed -i "s/whatever/${CERT_PASS}/g" /etc/raddb/mods-available/eap \
#SQL config
  && sed -i "s/<server>/${DB_HOST}/g" /etc/raddb/mods-available/sql \
  && sed -i "s/<port>/${DB_PORT}/g" /etc/raddb/mods-available/sql \
  && sed -i "s/<login>/${DB_USER}/g" /etc/raddb/mods-available/sql \
  && sed -i "s/<password>/${DB_PASSWORD}/g" /etc/raddb/mods-available/sql \
  && sed -i "s/<radius_db>/${DB_NAME}/g" /etc/raddb/mods-available/sql \
#Wifi Client config
  && sed -i "s/<client_ip>/${WIFI_CLIENT_IP}/g" /etc/raddb/clients.conf \
  && sed -i "s/<client_key>/${WIFI_CLIENT_KEY}/g" /etc/raddb/clients.conf \
  && sed -i "s/<client_ssid>/${WIFI_CLIENT_SSID}/g" /etc/raddb/clients.conf \
#Health Client config
  && sed -i "s/<health_net>/${HEALTH_CLIENT_NET}/g" /etc/raddb/clients.conf \
  && sed -i "s/<health_mask>/${HEALTH_CLIENT_MASK}/g" /etc/raddb/clients.conf \
  && sed -i "s/<health_key>/${HEALTH_CLIENT_KEY}/g" /etc/raddb/clients.conf \
  && sed -i "s/<health_port>/${HEALTH_CLIENT_PORT}/g" /etc/raddb/sites-available/inner-tunnel \
#Generate certificates
  && cd /etc/raddb/certs \
  #&& sed -i "s/2048/512/g" /etc/raddb/certs/bootstrap \
  && chmod -R 744 /etc/raddb/certs/bootstrap \
	&& ./bootstrap \
  && openssl ca -gencrl -keyfile ca.key -cert ca.pem -out crl.pem -config crl.cnf -passin pass:${CERT_PASS} \
  && cat ca.pem crl.pem > cacrl.pem
RUN mkdir -p /var/log/radius \
  && chown -R radius:radius /var/log/radius
EXPOSE 1812/udp 1813/udp

COPY --from=gobuilder /app/radius /bin/radius
COPY --from=gobuilder /app/healthy /bin/healthy
