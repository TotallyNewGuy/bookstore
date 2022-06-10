server:
	go run main.go

proto:
	rm -f pb_gen/book_gen/*.go
	rm -f pb_gen/order_gen/*.go
	rm -f pb_gen/user_info_gen/*.go
	protoc --proto_path=proto/book --go_out=pb_gen/book_gen --go_opt=paths=source_relative \
    --go-grpc_out=pb_gen/book_gen --go-grpc_opt=paths=source_relative proto/book/*.proto
	protoc --proto_path=proto/order --go_out=pb_gen/order_gen --go_opt=paths=source_relative \
    --go-grpc_out=pb_gen/order_gen --go-grpc_opt=paths=source_relative proto/order/*.proto
	protoc --proto_path=proto/user_info --go_out=pb_gen/user_info_gen --go_opt=paths=source_relative \
    --go-grpc_out=pb_gen/user_info_gen --go-grpc_opt=paths=source_relative proto/user_info/*.proto

evans:
	evans --host localhost --port 9090 -r repl

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server proto evans
