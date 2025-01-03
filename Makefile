postgres:
	docker run --name postgresgo --network simplebank_network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it postgresgo createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgresgo dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test

sqlc:
	sqlc generate

server:
	go run main.go
	
test:
	go test -v -cover ./... 