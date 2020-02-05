TOOL_VERSION_FILE=.tool_version
VERSION=`cat $(TOOL_VERSION_FILE)`

build:
	@go build -o bin/codacy-gorevive

unittest:
	@go test

run:
	@go run .

docgeneration:
	@go run ./docgenerator/.

docker:
	docker build --build-arg TOOL_VERSION=$(VERSION) -t codacy-gorevive .

all: 
	make unittest && make build && make docker
