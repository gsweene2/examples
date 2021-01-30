.PHONY: build clean deploy

build:
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/create todos/create.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/delete todos/delete.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/get todos/get.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/list todos/list.go
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/update todos/update.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose
