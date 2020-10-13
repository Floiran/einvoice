#!/bin/bash

set -ex

export API_SERVER_URL=http://localhost:8082
export PORT=8080
export CLIENT_BUILD_DIR=../client/build/
go run ./einvoice-web-app/server
popd
