default: build

code:
	coffee -c assets/application.coffee
	go generate

build:
	go install -ldflags "-X main.version=$(shell git describe --tags || git rev-parse --short HEAD || echo dev)"

upload:
	vault2env secret/aws/private -- aws s3 cp --acl=public-read ~/.local/share/ed-fast-travel/dump_v3.bin.gz s3://assets.luzifer.io/
