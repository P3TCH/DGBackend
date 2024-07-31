build:
	@go build -o bin/DG cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/DG
