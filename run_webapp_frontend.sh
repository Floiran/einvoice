#!/bin/bash

set -ex

pushd einvoice-web-app/client
export API_SERVER_URL=http://localhost:8082
npm start
popd
