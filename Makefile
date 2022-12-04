postgres:
	docker run --name postgres-local -p 5432:5432 -e=POSTGRES_USER=root -e=POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres-local createdb --username=root --owner=root apple_store

dropdb:
	docker exec -it postgres-local dropdb apple_store

migrateup:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/apple_store?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/apple_store?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen --destination db/mock/store.go --package mockdb  github.com/alekseiapa/apple_store/db/sqlc Store


.PHONY: database test api mock

# phony targets
database : postgres createdb dropdb migrateup migratedown sqlc
api: server