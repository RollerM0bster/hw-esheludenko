#!/bin/bash
set -e

DB_URL="postgres://postgres:postgres@localhost:5432/calendardb?sslmode=disable&search_path=eventstorage"
migrate -path ./migrations -database "$DB_URL" up
