run-db:
	 docker run --name=kokos -e POSTGRES_PASSWORD='12345' -p 5432:5432 -d --rm postgres

migrations-up:
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' up

migrations-down:
	migrate -path ./schema -database 'postgres://postgres:12345@localhost:5432/postgres?sslmode=disable' down

run-server:
	go run cmd/main.go