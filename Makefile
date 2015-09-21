all: build/pbcopy build/pbpaste

build/pbcopy:	pbcopy.go common.go;
	go build -o build/pbcopy pbcopy.go common.go

build/pbpaste:	pbpaste.go common.go;
	go build -o build/pbpaste pbpaste.go common.go

clean:
	rm -rf ./build
