createdb:
	createdb user-golang

dropdb:
	dropdb user-golang

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5433/postgres?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5433/postgres?sslmode=disable" --verbose down
