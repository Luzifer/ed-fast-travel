default: build

build:
	go build -ldflags "-X main.version=$(git describe --tags || git rev-parse --short HEAD || echo dev)"
