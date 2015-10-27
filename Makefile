VERSION = $(shell cat VERSION)

all: build/pbcopy build/pbpaste build/pbserve build/rpb.tgz

build/pbcopy:	pbcopy.go common.go constants.go;
	go build -o build/pbcopy -ldflags "-X main.Version ${VERSION}" pbcopy.go common.go constants.go

build/pbpaste:	pbpaste.go common.go constants.go;
	go build -o build/pbpaste -ldflags "-X main.Version ${VERSION}" pbpaste.go common.go constants.go

build/pbserve:	pbserve.go constants.go;
	go build -o build/pbserve -ldflags "-X main.Version ${VERSION}" pbserve.go constants.go

build/rpb.tgz: build/pbpaste build/pbcopy
	tar czvf build/rpb.tgz -C build/ pbcopy pbpaste

clean:
	rm -rf ./build
