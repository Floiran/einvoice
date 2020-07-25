#!/bin/bash

set -ex

psql postgres -h 127.0.0.1 -d einvoice -f setup_db.sql
echo "setup_db.sql executed"