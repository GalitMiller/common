.PHONY: test build clean deploy run gomodgen

test:
	export GOPATH="${PWD}"
	export GO111MODULE=on
	go test fetch/*

build: gomodgen clean test
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/fetch fetch/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: build
	sls deploy --verbose

run:
	sls invoke -f fetch

gomodgen:
	chmod u+x ../gomod.sh
	../gomod.sh
