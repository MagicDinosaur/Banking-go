postgres:
	docker run --name gobanking -e POSTGRES_USER=root -e POSTGRES_PASSWORD=1234 -p 5432:5432 -d postgres:12

createdb:
	docker exec -it gobanking createdb --username=root --owner=root gobanking

dropdb:
	docker exec -it postgres dropdb gobanking

migrateup:
	migrate -path db/migrations -database "postgresql://root:1234@localhost:5432/gobanking?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:1234@localhost:5432/gobanking?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc