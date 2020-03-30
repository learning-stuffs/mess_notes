#!/bin/bash
#openssl genrsa -out app.rsa 4096
#openssl rsa -in app.rsa -pubout > app.rsa.pub

openssl req -x509 -newkey rsa:4096 \
-keyout server-key.pem \
-out server-cert.pem \
-days 365 -nodes -subj '/CN=localhost'