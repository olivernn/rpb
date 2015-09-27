all: build/pbcopy build/pbpaste
VERSION = $(shell cat VERSION)

build/pbcopy:	pbcopy.go common.go;
	go build -o build/pbcopy -ldflags "-X main.version=${VERSION}" pbcopy.go common.go

build/pbpaste:	pbpaste.go common.go;
	go build -o build/pbpaste -ldflags "-X main.version=${VERSION}" pbpaste.go common.go

clean:
	rm -rf ./build
