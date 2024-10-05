run:
	@go run cmd/server/main.go

build:
	@go build -o main cmd/server/main.go

start:
	@./main