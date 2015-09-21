all: build/pbcopy build/pbpaste

build/pbcopy:	pbcopy/pbcopy.go;
	go build -o build/pbcopy pbcopy/pbcopy.go

build/pbpaste:	pbpaste/pbpaste.go;
	go build -o build/pbpaste pbpaste/pbpaste.go

clean:
	rm -rf ./build
