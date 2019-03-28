#!/bin/sh
# this will create a self signed certificate and the matching private key in order to start in tls mode.
# this is meant to run the app with encryption enabled but verification disabled. If you want to reate verifiable
# certificates, you will need to use a different technique
NAME=xxx
openssl req -newkey rsa:2048 -nodes -keyout ${NAME}.key -x509 -days 3650 -out ${NAME}.crt -subj "/C=DE/ST=Berlin/L=Berlin/O=kesselFAKTUR/OU=Org/CN=${NAME}"
