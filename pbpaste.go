package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	port := flag.Int("port", 2225, "remote pbpaste port")
	flag.Parse()
	pboard, err := OpenRead(*port)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, pboard)
}
