BINARY_NAME=golang_web

build:
	$(MAKE) buildFrontend
	$(MAKE) runTempl
	GOARCH=amd64 GOOS=darwin go build -o dist/${BINARY_NAME}-darwin app/*.go
	GOARCH=amd64 GOOS=linux go build -o dist/${BINARY_NAME}-linux app/*.go
	GOARCH=amd64 GOOS=windows go build -o dist/${BINARY_NAME}-windows app/*.go
	GOARCH=arm64 GOOS=darwin go build -o dist/${BINARY_NAME}-arm64-darwin app/*.go
	GOARCH=arm64 GOOS=linux go build -o dist/${BINARY_NAME}-arm64-linux app/*.go
	GOARCH=arm64 GOOS=windows go build -o dist/${BINARY_NAME}-arm64-windows app/*.go

dockerBuild:
	$(MAKE) buildFrontend
	$(MAKE) runTempl
	go build -o /go/bin/app app/*.go

watchTempl:
	templ generate -watch

runTempl:
	templ generate

buildFrontend:
	./node_modules/.bin/tailwindcss build -o public/styles.css

devFrontend:
	./node_modules/.bin/tailwindcss -i app/resources/styles.css -o public/styles.css --watch

devServer:
	air

watch:
	make -j3 runTempl devFrontend devServer

hello:
	echo "Hello, World!"

buildDockerImage:
	docker build -t golang-web .

runDocker:
	docker compose up --build -d