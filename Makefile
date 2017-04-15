default: build

ci:
	curl -sSLo golang.sh https://raw.githubusercontent.com/Luzifer/github-publish/master/golang.sh
	bash golang.sh

code:
	coffee -c assets/application.coffee
	go generate

build:
	go build -ldflags "-X main.version=$(shell git describe --tags || git rev-parse --short HEAD || echo dev)"

install:
	go install -ldflags "-X main.version=$(shell git describe --tags || git rev-parse --short HEAD || echo dev)"

update-db-dump:
	ed-fast-travel --generate-database --data-path=/tmp
	gzip -c /tmp/dump_v3.bin > /tmp/dump_v3.bin.gz
	vault2env aws/sts/ed-fast-travel \
		-t access_key=AWS_ACCESS_KEY_ID \
		-t secret_key=AWS_SECRET_ACCESS_KEY \
		-t security_token=AWS_SECURITY_TOKEN \
		-- \
		aws s3 cp --acl=public-read /tmp/dump_v3.bin.gz s3://assets.luzifer.io/
