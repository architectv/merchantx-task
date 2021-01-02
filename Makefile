build:
	docker-compose build app

run:
	docker-compose up app

# WARNING: need to create database 'postgres_test' in postgres localhost
run_test:
	go test ./... -coverprofile cover.out
	go test -tags=e2e

migrateup:
	migrate -path ./scripts -database 'postgres://postgres:1234@0.0.0.0:5436/postgres?sslmode=disable' up

migratedown:
	migrate -path ./scripts -database 'postgres://postgres:1234@0.0.0.0:5436/postgres?sslmode=disable' down