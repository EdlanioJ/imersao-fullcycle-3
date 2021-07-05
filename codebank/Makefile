start:
	go run main.go

build:
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o codebank ./main.go

.PHONY: start build