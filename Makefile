all: build/pbcopy build/pbpaste

build/pbcopy:	pbcopy.go rpb.go;
	go build -o build/pbcopy pbcopy.go rpb.go

build/pbpaste:	pbpaste.go rpb.go;
	go build -o build/pbpaste pbpaste.go rpb.go

clean:
	rm -rf ./build
