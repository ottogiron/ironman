# Meta info
NAME := ironman
VERSION := 0.1.0
MAINTAINER := "Otto Giron <ottog2486@gmail.com"
SOURCE_URL := https://github.com/ottogiron/ironman.git
DATE := $(shell date -u +%Y%m%d.%H%M%S)
COMMIT_ID := $(shell git rev-parse --short HEAD)
GIT_REPO := $(shell git config --get remote.origin.url)
# Go tools flags
LD_FLAGS := -X 	github.com/ottogiron/ironman/cmd.buildVersion=$(VERSION)
LD_FLAGS += -X github.com/ottogiron/ironman/cmd.buildCommit=$(COMMIT_ID)
LD_FLAGS += -X github.com/ottogiron/ironman/cmd.buildDate=$(DATE)
EXTRA_BUILD_VARS := CGO_ENABLED=0 GOARCH=amd64
SOURCE_DIRS := $(shell go list ./... | grep -v /vendor/)


all: test package-linux package-darwin

build-release: container

lint:
	@go fmt $(SOURCE_DIRS)
	@go vet $(SOURCE_DIRS)

test: install_dependencies lint
	 @go test -v $(SOURCE_DIRS) -cover -bench . -race 

install_dependencies: 
	glide install

cover: 
	gocov test $(SOURCE_DIRS) | gocov-html > coverage.html && open coverage.html
	

image: binaries
	docker-flow build -f docker/Dockerfile


binaries: binary-darwin binary-linux

binary-darwin:
	@-rm -rf build/dist/darwin
	@-mkdir -p build/dist/darwin
	GOOS=darwin $(EXTRA_BUILD_VARS) go build -ldflags "$(LD_FLAGS)" -o build/dist/darwin/$(NAME)

binary-linux:
	@-rm -rf build/dist/linux
	@-mkdir -p build/dist/linux
	GOOS=linux $(EXTRA_BUILD_VARS) go build -ldflags "$(LD_FLAGS)" -o build/dist/linux/$(NAME)


package-darwin: binary-darwin
	@tar -czf build/dist/ironman.darwin-amd64.tar.gz -C build/dist/darwin ironman


package-linux: binary-linux
	@tar -czf build/dist/ironman.linux-amd64.tar.gz -C build/dist/linux ironman