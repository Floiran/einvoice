#!/bin/bash

set -ex

psql postgres -h 127.0.0.1 -d einvoice
