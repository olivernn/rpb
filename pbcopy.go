package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	port := flag.Int("port", 2224, "remote pbcopy port")
	flag.Parse()
	pboard, err := OpenWrite(*port)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(pboard, os.Stdin)
}
