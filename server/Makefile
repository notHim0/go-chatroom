postgresinit: 
	docker run --name postgresdb -p 5000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:alpine

postgres: 
	docker exec -it postgresdb psql

createdb: 
	docker exec -it postgresdb createdb --username=root --owner=root go-chat

dropdb:
	docker exec -it postgresdb dropdb go-chat
migrateup:
	migrate -path db/migrations/ -database "postgresql://root:password@localhost:5000/go-chat?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations/ -database "postgresql://root:password@localhost:5000/go-chat?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb