# GO backend Meteo (Web Service)

Pet project weather station based on esp32 controller. Frontend is designed using quasar. The backend is developed on golang and built according to the microservice scheme using docker compose.

## Prepare install

**Generate Certificate**
~cd /tmp
~git clone https://github.com/vivask/Meteo.Backend.git
~cd Meteo.Backend/certs
Editing the [ req_distinguished_name ] section in the ca.cnf, client.cnf, server.cnf
~./makecerts.sh

**Build frontend**
~cd /tmp
~git clone https://github.com/vivask/Meteo.Ui.git
~cd Meteo.Ui
~npm install
~quasar build
~cp -r ./dist/spa/* ../Meteo.Backend/ui/

**ESP32 install**
Controller firmware is described here https://github.com/vivask/Meteo.ESP32

## Install
~cd ../Meteo.Backend
~docker compose -f release.yaml up --build
