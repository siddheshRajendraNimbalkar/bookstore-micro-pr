.PHONY: 

postgres:
	docker run --name bookstore-order -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

createdb:
	docker exec -it bookstore-order createdb --username=root --owner=root bookstore-orderDB

dropdb:
	docker exec -it bookstore-order dropdb --username=root bookstore-orderDB

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/bookstore-orderDB?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/bookstore-orderDB?sslmode=disable" -verbose down

sqlc:
	sqlc generate

generateProto:
	rm -rf pb
	mkdir -p pb
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/*.proto