#!/bin/sh

#docker exec -it backend-xu4-1 sh

echo "Service 'All': Status"
rc-status -a
echo "Service 'RSyslog': Starting ..."
rc-service rsyslog start

echo "Service 'Sshd': Starting ..."
rc-update add sshd
service sshd start

echo "Service 'Samba': Starting ..."
rc-update add samba
rc-service samba start

echo "Service 'Transmission': Starting ..."
rc-update add transmission-daemon
rc-service transmission-daemon start

echo "Service 'Haveged': Starting ..."
rc-update add haveged default
rc-service haveged start

cd /etc/raddb/certs
export CA_PASS=$(cat /app/docker/secrets/ca_password.txt)
openssl ca -gencrl -keyfile ca.key -cert ca.pem -out crl.pem -config crl.cnf -passin pass:$CA_PASS
cat ca.pem crl.pem > cacrl.pem

echo "Service 'FreeRadius': Starting ..."
rc-update add radiusd default
rc-service radiusd start

