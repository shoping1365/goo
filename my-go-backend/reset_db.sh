#!/bin/bash
# Reset PostgreSQL database schema (drop and recreate public schema)
# Update the following variables as needed
PGUSER="your_user"
PGDATABASE="your_db"

psql -U "$PGUSER" -d "$PGDATABASE" -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;" 