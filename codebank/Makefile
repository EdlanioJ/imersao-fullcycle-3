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

mock:
	mockery --output "./domain/mocks" --dir "./" --all

.PHONY: start build create-migration migrate-up migrate-down mock
