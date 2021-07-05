DATABASE="postgresql://postgres:root@db:5432/codebank?sslmode=disable"

start:
	go run main.go

build:
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o codebank ./main.go

create-migration:
	migrate create -ext sql -dir infrastructure/db/migration -seq init_schema

migrate-up:
	migrate -path infrastructure/db/migration -database ${DATABASE} -verbose up

migrate-down:
	migrate -path infrastructure/db/migration -database ${DATABASE} -verbose down

.PHONY: start build create-migration migrate-up migrate-down