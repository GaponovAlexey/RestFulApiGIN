.PHONY: build

up:
	migrate -path ./shema -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable up


create2:
	docker run --name pg -d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres  -p 5432:5432 -it postgres:14.1-alpine
