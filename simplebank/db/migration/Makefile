postgres:
	docker run --name postgres12 \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=18845383 \
	-e POSTGRES_DB=simple_bank \
	-p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb -U postgres -O postgres simple_bank

dropdb:
	docker exec -it postgres12 dropdb -U postgres simple_bank

migrateup:
	migrate -path . -database "postgres://postgres:18845383@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path . -database "postgres://postgres:18845383@localhost:5432/simple_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown