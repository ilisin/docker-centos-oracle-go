#!/bin/sh
tar xvf go1.5.1.linux-amd64.tar.gz
docker build -t 192.168.10.40:5000/linux-go-oracle ./
rm -rf go
