postgres:
	docker run --name postgres-local --network applestore-net -p 5432:5432 -e=POSTGRES_USER=root -e=POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres-local createdb --username=root --owner=root apple_store

dropdb:
	docker exec -it postgres-local dropdb apple_store

migrateup:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/apple_store?sslmode=disable" --verbose up

migrateup1ver:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/apple_store?sslmode=disable" --verbose up 1

migratedown:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/apple_store?sslmode=disable" --verbose down

migratedown1ver:
	migrate --path db/migration --database "postgresql://root:secret@localhost:5432/apple_store?sslmode=disable" --verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

docker_network:
	docker network create app_net

docker_compose:
	docker-compose up --build

mock:
	mockgen --destination db/mock/store.go --package mockdb  github.com/alekseiapa/apple_store/db/sqlc Store



.PHONY: database test api mock docker

# phony targets
database : postgres createdb dropdb migrateup migratedown sqlc migrateup1ver migratedown1ver
api: server
docker: docker_compose docker_network