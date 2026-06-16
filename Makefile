.PHONY: run build tidy clean

run:
	go run ./cmd/main.go

build:
	go build -o bin/app ./cmd/main.go

tidy:
	go mod tidy

clean:
	rm -f bin/app products.db
