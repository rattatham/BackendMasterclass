postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	/usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/cq/tx2yxt4d4v52wfp7rn0g6n100000gn/T/vscode-gofvTlOi/go-code-cover github.com/techschool/db/sqlc

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server

