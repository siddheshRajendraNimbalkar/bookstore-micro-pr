.PHONY: 

postgres:
	docker run --name bookstore-product -p 5431:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

createdb:
	docker exec -it bookstore-product createdb --username=root --owner=root bookstore-productDB

dropdb:
	docker exec -it bookstore-product dropdb --username=root bookstore-productDB

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5431/bookstore-productDB?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5431/bookstore-productDB?sslmode=disable" -verbose down

sqlc:
	sqlc generate

generateProto:
	rm -rf pb
	mkdir -p pb
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
		--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		proto/*.proto