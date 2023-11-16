build:
	@go build -o ./bin/goqu cmd/goqu/main.go

run: build
	@./bin/goqu

test:
	@go test -v ./...