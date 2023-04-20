#!/bin/sh

PGPASSWORD=taylor psql --username=taylor --dbname=taylor < pg/migrate.sql
