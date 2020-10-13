.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o bin/create-product aws/create-product/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/create-workspace aws/create-workspace/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get-product-with-id aws/get-product-with-id/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get-products aws/get-products/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/get-workspace-ownerships aws/get-workspace-ownerships/main.go

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy-devingen: clean build
	serverless deploy --stage prod --region eu-central-1 --verbose

teardown-devingen: clean
	serverless remove --stage prod --region eu-central-1 --verbose

deploy-devingen-dev: clean build
	serverless deploy --stage dev --region ca-central-1 --verbose

teardown-devingen-dev: clean
	serverless remove --stage dev --region ca-central-1 --verbose
