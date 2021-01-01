build:
	docker-compose build app

run:
	docker-compose up app

test:
	go test -v ./...

migrateup:
	migrate -path ./scripts -database 'postgres://postgres:1234@0.0.0.0:5436/postgres?sslmode=disable' up

migratedown:
	migrate -path ./scripts -database 'postgres://postgres:1234@0.0.0.0:5436/postgres?sslmode=disable' down