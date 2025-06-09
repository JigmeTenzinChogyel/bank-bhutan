DB_URL=postgresql://root:secret@localhost:5432/bank_bhutan?sslmode=disable

postgres:
	docker run --name postgres17 --network bank-netwrok -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root bank_bhutan

dropdb:
	docker exec -it postgres17 dropdb --username=root --owner=root bank_bhutan

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down


migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

generate:
	sqlc generate
	
test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/JigmeTenzinChogyel/bank-bhutan/db/sqlc Store

proto:
	rm -rf proto/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock migrateup1 migratedown1 proto