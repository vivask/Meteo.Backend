#!/bin/bash

rm -f ./*.crt ./*.key
rm -f ./crt/* ./csr/* ./key/* ./newcerts/*

openssl genrsa -aes256 -out key/ca.key 2048
openssl req -config ca.cnf -key key/ca.key -new -x509 -days 3650 -sha256 -extensions ca_cert -out crt/ca.crt

openssl genrsa -out key/server.key 2048
openssl req -config server.cnf -key key/server.key -new -sha256 -out csr/server.csr
openssl ca -config server.cnf -extensions server_cert -days 3650 -notext -md sha256 -in csr/server.csr -out crt/server.crt

openssl genrsa -out key/client.key 2048
openssl req -config client.cnf -key key/client.key -new -sha256 -out csr/client.csr
openssl ca -config client.cnf -extensions usr_cert -days 3650 -notext -md sha256 -in csr/client.csr -out crt/client.crt

cp ./crt/* ./
cp ./key/* ./


#cd ./firmware
#openssl s_server -accept 8070 -WWW -CAfile ../certs/crt/ca.crt -key ../certs/key/server.key -cert ../certs/crt/server.crt -tls1_2 -state -Verify 3

#openssl x509 -in ca.crt -text

