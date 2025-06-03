build:
	go build -o terradrift main.go

local_test: build
	sudo chmod +x terradrift
	sudo mv terradrift /usr/local/bin/terradrift 

test:
	go test -v ./...