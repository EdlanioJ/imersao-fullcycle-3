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

gen:
	protoc --proto_path=application/grpc application/grpc/protofile/*.proto --go_out=application/ --go-grpc_out=application/

.PHONY: start build create-migration migrate-up migrate-down mock gen
