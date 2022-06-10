#! /bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "postgresql://root:AppleMan922@postgres:5432/book_store?sslmode=disable" -verbose up

echo "start the app"
exec "$@"