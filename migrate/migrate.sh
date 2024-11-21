#!/bin/bash

#export PGPASSWORD=admin
#export PGUSER=admin
#export PGHOST=127.0.0.1
#export PGPORT=5432
#export PGDATABASE=postgres

psql -w -f migrate/create_users.sql