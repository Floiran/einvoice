#!/bin/bash

set -ex

# Check if valid service name was provided
if ! [[ "$1" =~ ^(apiserver|authproxy|einvoice-web-app/server)$ ]]; then
  echo Service "$1" does not exist.
  exit 1
fi

# Move to project root directory
cd "$(dirname "$0")"/..
# Load env variables
if [ ! -f "$1"/.env ]; then
    echo "$1"/.env not found. Create it and fill with env variables.
    exit 1
fi
# shellcheck disable=SC2046
export $(xargs < "$1"/.env)

# Run go service
go run ./"$1"
