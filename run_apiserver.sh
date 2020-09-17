#!/bin/bash

set -ex

#./db.sh --drop --create --setup
#redis-server

export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=a
export DB_NAME=einvoice
export SLOW_STORAGE_TYPE=local
export LOCAL_STORAGE_BASE_PATH=/home/filip/einvoiceStorage
export D16B_XSD_PATH=xml/d16b/xsd
export UBL21_XSD_PATH=xml/ubl21/xsd
export PORT=8081
go run ./apiserver
