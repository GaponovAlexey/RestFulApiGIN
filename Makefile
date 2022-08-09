.PHONY: build

up:
	migrate -path ./shema -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable up


create2:
	docker compose up -d
