run: build
	@bin/cmd/app

build:
	go build -o bin/cmd/app cmd/maze/main.go