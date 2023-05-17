postgres:
	docker run --name graphqldb -p 5432:5432  -e POSTGRES_USER=graphql  -e POSTGRES_PASSWORD=graphql -d postgres:latest

createdb:
	docker exec -it graphqldb createdb --username=graphql --owner=graphql graphql_data

migratecreate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	 migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

dropdb:
	docker exec -it scalefocusdb dropdb $(DB_NAME)

migratedown:
	migrate -path db/migration -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

.PHONY: postgres createdb createtestdb dropdb migrateup migratedown migratecreate