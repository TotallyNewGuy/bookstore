postgres:
	docker run --name postgres142 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=AppleMan922 -d postgres:14.2-alpine

createdb:
	docker exec -it postgres142 createdb --username=root --owner=root book_store

dropdb:
	docker exec -it postgres142 dropdb book_store

migratedown:
	migrate -path db/migration -database "postgresql://root:AppleMan922@localhost:5432/book_store?sslmode=disable" -verbose down

migrateup:
	migrate -path db/migration -database "postgresql://root:AppleMan922@localhost:5432/book_store?sslmode=disable" -verbose up

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

proto:
	rm -f pb_gen/order_gen/*.go
	protoc --proto_path=proto --go_out=pb_gen/order_gen --go_opt=paths=source_relative \
    --go-grpc_out=pb_gen/order_gen --go-grpc_opt=paths=source_relative \
    proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server proto evans