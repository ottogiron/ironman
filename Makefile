# Meta info
NAME := ironman
VERSION := v0.4.0
MAINTAINER := "Otto Giron <ottog2486@gmail.com"
SOURCE_URL := https://github.com/ironman-project/ironman.git
DATE := $(shell date -u +%Y%m%d.%H%M%S)
COMMIT_ID := $(shell git rev-parse --short HEAD)
GIT_REPO := $(shell git config --get remote.origin.url)
# Go tools flags
LD_FLAGS := -X github.com/ironman-project/ironman/cmd.buildVersion=$(VERSION)
LD_FLAGS += -X github.com/ironman-project/ironman/cmd.buildCommit=$(COMMIT_ID)
LD_FLAGS += -X github.com/ironman-project/ironman/cmd.buildDate=$(DATE)
EXTRA_BUILD_VARS := CGO_ENABLED=0 GOARCH=amd64
SOURCE_DIRS := ./...
SUBDIRS = acceptance

all: test package-linux package-darwin package-windows acceptance

build-release: container

lint:
	go fmt $(SOURCE_DIRS)
	go vet $(SOURCE_DIRS)

test:  lint
	 go test -v $(SOURCE_DIRS) -cover -bench . -race

cover: 
	gocov test $(SOURCE_DIRS) | gocov-html > coverage.html && open coverage.html
	
binaries: binary-darwin binary-linux

binary-darwin:
	@-rm -rf build/dist/darwin
	@-mkdir -p build/dist/darwin
	GOOS=darwin $(EXTRA_BUILD_VARS) go build -ldflags "$(LD_FLAGS)" -o build/dist/darwin/$(NAME)

binary-linux:
	@-rm -rf build/dist/linux
	@-mkdir -p build/dist/linux
	GOOS=linux $(EXTRA_BUILD_VARS) go build -ldflags "$(LD_FLAGS)" -o build/dist/linux/$(NAME)

binary-windows:
	@-rm -rf build/dist/windows
	@-mkdir -p build/dist/windows
	GOOS=windows $(EXTRA_BUILD_VARS) go build -ldflags "$(LD_FLAGS)" -o build/dist/windows/$(NAME).exe


package-darwin: binary-darwin
	@tar -czf build/dist/ironman.darwin-amd64.tar.gz -C build/dist/darwin ironman


package-linux: binary-linux
	@tar -czf build/dist/ironman.linux-amd64.tar.gz -C build/dist/linux ironman

package-windows: binary-windows
	@tar -czf build/dist/ironman.windows-amd64.tar.gz -C build/dist/windows ironman.exe 

.PHONY: $(SUBDIRS)
