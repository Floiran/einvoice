#!/bin/bash

set -ex

#./db.sh --drop --create --setup
#redis-server

export APISERVER_URL=http://localhost:8081
export REDIS_URL=localhost:6379
export PORT=8082
go run ./authproxy
