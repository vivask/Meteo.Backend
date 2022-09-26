FROM golang:alpine
RUN apk --no-cache add git gcc make mc shadow
RUN adduser --disabled-password odroid
RUN apk --no-cache add samba \
  && mkdir -p /media/samba/storage
ADD ./docker/secrets/user_name.txt /var/tmp/user_name
ADD ./docker/secrets/user_password.txt /var/tmp/user_password
ADD ./docker/secrets/samba_password.txt /var/tmp/samba_password
RUN export USER_NAME=$(cat /var/tmp/user_name) \
  && export USER_PASS=$(cat /var/tmp/user_password) \
  && export SAMBA_PASS=$(cat /var/tmp/samba_password) \
	&& rm -f /var/tmp/user_password /var/tmp/samba_password \
	&& usermod --password $USER_PASS $USER_NAME \
  && echo -ne "$SAMBA_PASS\n$SAMBA_PASS\n" | smbpasswd -a -s $USER_NAME
RUN apk update \
  && apk add --no-cache rsyslog openssh-server openrc rsync \
  && mkdir -p /run/openrc \
  && touch /run/openrc/softlevel \
  && echo "Port 2222" >> /etc/ssh/sshd_config\
  && echo "AddressFamily inet"  >> /etc/ssh/sshd_config\
  && echo "ListenAddress 0.0.0.0"  >> /etc/ssh/sshd_config
ADD ./docker/xu4/smb.conf /etc/samba/smb.conf
RUN apk --no-cache add transmission-daemon \
  && mkdir -p /etc/transmission-daemon \
	&& echo "TRANSMISSION_OPTIONS=\"--encryption-preferred\"" > /etc/conf.d/transmission-daemon \
	&& echo "runas_user=transmission" >> /etc/conf.d/transmission-daemon \
	&& echo "logfile=/var/log/transmission/transmission.log"  >> /etc/conf.d/transmission-daemon \
  && mkdir -p /var/lib/transmission/config \
  && chown -R transmission:transmission /var/lib/transmission/config
ADD ./docker/xu4/settings.json /etc/transmission-daemon/settings.json
RUN chown -R transmission:transmission /etc/transmission-daemon/settings.json \
	&& ln -s /etc/transmission-daemon/settings.json /var/lib/transmission/config/settings.json
RUN apk --no-cache add freeradius freeradius-eap freeradius-utils freeradius-postgresql openssl haveged
ADD ./docker/xu4/radius/certs/ca.cnf /etc/raddb/certs/ca.cnf
ADD ./docker/xu4/radius/certs/client.cnf /etc/raddb/certs/client.cnf
ADD ./docker/xu4/radius/certs/inner-server.cnf /etc/raddb/certs/inner-server.cnf
ADD ./docker/xu4/radius/certs/server.cnf /etc/raddb/certs/server.cnf
ADD ./docker/xu4/radius/certs/crl.cnf /etc/raddb/certs/crl.cnf
ADD ./docker/xu4/radius/clients.conf /etc/raddb/clients.conf
ADD ./docker/xu4/radius/mods-available/eap /etc/raddb/mods-available/eap
ADD ./docker/xu4/radius/mods-available/sql /etc/raddb/mods-available/sql
ADD ./docker/xu4/radius/sites-available/default /etc/raddb/sites-available/default
ADD ./docker/xu4/radius/sites-available/inner-tunnel /etc/raddb/sites-available/inner-tunnel
ADD ./docker/xu4/radius/sites-available/tls /etc/raddb/sites-available/tls
RUN cd /etc/raddb/certs \
  && ./bootstrap


RUN mkdir /app
WORKDIR /app
# RUN go get -u honnef.co/go/tools
ADD go.mod .
ADD go.sum .
RUN go mod download

ENV GIN_MODE=debug

#ENTRYPOINT ["sh","-c", "rc-status -a; rc-service sshd start; crond -f"]
