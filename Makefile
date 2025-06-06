postgres:
	docker run --name postgres17 --network bank-netwrok -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root bank_bhutan

dropdb:
	docker exec -it postgres17 dropdb --username=root --owner=root bank_bhutan

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_bhutan?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_bhutan?sslmode=disable" -verbose up 1


migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_bhutan?sslmode=disable" -verbose down


migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank_bhutan?sslmode=disable" -verbose down 1

generate:
	sqlc generate
	
test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/JigmeTenzinChogyel/bank-bhutan/db/sqlc Store
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock migrateup1 migratedown1