createdb:
	docker exec -it docker-user-golang-database-1 createdb --username=postgres --owner=postgres user_golang

dropdb:
	docker exec -it docker-user-golang-database-1 dropdb --username=postgres user_golang

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5433/user_golang?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:password@localhost:5433/user_golang?sslmode=disable" --verbose down
