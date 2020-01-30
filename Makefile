TOOL_VERSION_FILE=.version
VERSION=`cat $(TOOL_VERSION_FILE)`

build:
	@go build -o bin/codacy-gorevive

test:
	@go test

run:
	@go run .

docgeneration:
	@go run ./docgenerator/.

docker:
	GOOS=linux GOARCH=amd64 go build -o bin/codacy-gorevive
	docker build --build-arg TOOL_VERSION=$(VERSION) -t codacy-gorevive .

all:
	test build
