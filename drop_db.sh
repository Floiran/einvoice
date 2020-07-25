#!/bin/bash

set -ex

psql postgres -h 127.0.0.1 -d postgres -f drop_db.sql
echo "drop_db.sql executed"