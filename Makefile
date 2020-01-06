
build:
	@go build -o bin/codacy-gorevive

test:
	@go test

run:
	@go run .

docker:
	GOOS=linux GOARCH=amd64 go build -o bin/codacy-gorevive
	docker build -t codacy-gorevive .

all:
	test build
