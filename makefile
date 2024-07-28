migrateup:
	migrate -path db/migration -database "postgresql://postgres:pass@localhost:5431/postgres?sslmode=disable" -verbose up

migrateforce:
	migrate -path db/migration -database "postgresql://postgres:pass@localhost:5431/postgres?sslmode=disable" force 1

migratedown: migrateforce
	migrate -path db/migration -database "postgresql://postgres:pass@localhost:5431/postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run ./cmd/main.go

.PHONY: migrateup migratedown migrateforce sqlc server

all: test
