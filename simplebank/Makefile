postgres:
	docker run --name postgres12 -p 5432:5432 \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=18845383 \
	-e POSTGRES_DB=simple_bank \
	-d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb -U postgres simple_bank

dropdb:
	docker exec -it postgres12 dropdb -U postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgres://postgres:18845383@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://postgres:18845383@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc


# To run the Postgres troubleshooting port conflict with docker 

# docker start postgres12
# docker stop postgres12

#  In power shell to test the port connection
#  Test-NetConnection -ComputerName localhost -Port 5432

# netstat -ano | findstr :5432

# Stop-Service postgresql-x64-12
# Stop-Service postgresql-x64-13
# Stop-Service postgresql-x64-15
# Stop-Service postgresql-x64-17

# Stop-Service postgresql*

#   end of the power shell commands # 