VERSION = $(shell cat VERSION)

all: build/pbcopy build/pbpaste build/rpb.tgz

build/pbcopy:	pbcopy.go common.go;
	go build -o build/pbcopy -ldflags "-X main.version=${VERSION}" pbcopy.go common.go

build/pbpaste:	pbpaste.go common.go;
	go build -o build/pbpaste -ldflags "-X main.version=${VERSION}" pbpaste.go common.go

build/pbserve:	pbserve.go;
	go build -o build/pbserve pbserve.go

build/rpb.tgz: build/pbpaste build/pbcopy
	tar czvf build/rpb.tgz -C build/ pbcopy pbpaste

clean:
	rm -rf ./build
