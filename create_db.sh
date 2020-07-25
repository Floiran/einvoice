#!/bin/bash

set -ex

psql postgres -h 127.0.0.1 -d postgres -f create_db.sql
echo "create_db.sql executed"
