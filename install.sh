#!/bin/sh

if [ $USER != "root" ]; then
  echo "Eres root"
  exit 1
fi

systemctl stop proxy-logger > /dev/null
cp bin/proxy-logger /usr/bin/proxy-logger
mkdir -p /etc/proxy-logger
cp config.yml /etc/proxy-logger/config.yml
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