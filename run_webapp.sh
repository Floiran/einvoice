#!/bin/bash

set -ex

#./db.sh --drop --create --setup
#redis-server

pushd einvoice-web-app/client
export API_SERVER_URL=http://localhost:8082
npm start
popd
