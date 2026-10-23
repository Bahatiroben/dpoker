build:
	go build -o bin/amakarita main.go

run: build
	./bin/amakarita

test: go test ./...