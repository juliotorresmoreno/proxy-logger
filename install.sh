#!/bin/bash

if [ $USER != "root" ]; then
  echo "No eres root"
  exit 1
fi

systemctl stop proxy-logger > /dev/null
cp bin/proxy-logger /usr/bin/proxy-logger
mkdir -p /etc/proxy-logger
cp config.yml /etc/proxy-logger &> /dev/null
cp config.yml.example /etc/proxy-logger &> /dev/null
cp proxy-logger.service /etc/systemd/system
adduser proxy-logger \
    --gecos "" \
    --system \
    --no-create-home \
    --disabled-password \
    --disabled-login > /dev/null
systemctl daemon-reload
systemctl start proxy-logger
systemctl status proxy-logger
