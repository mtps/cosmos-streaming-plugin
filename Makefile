all: plugin

plugin:
	go build -buildmode=plugin ./...
