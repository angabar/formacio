createdb:
	PGPASSWORD=admin123 psql -U postgres -c 'create database simple_bank_db;'

dropdb:
	PGPASSWORD=admin123 psql -U postgres -c 'drop database simple_bank_db;'

migrateup:
	migrate -path db/migration -database "postgresql://postgres:admin123@localhost:5432/simple_bank_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:admin123@localhost:5432/simple_bank_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY:
	createdb dropdb migrateup migratedown sqlc