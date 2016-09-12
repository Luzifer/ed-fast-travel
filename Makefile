default: build

code:
	coffee -c assets/application.coffee
	go generate

build:
	go install -ldflags "-X main.version=$(shell git describe --tags || git rev-parse --short HEAD || echo dev)"
