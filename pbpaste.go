package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var version string

func main() {
	port := flag.Int("port", 2225, "remote pbpaste port")
	showVersion := flag.Bool("v", false, "display version")

	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	pboard, err := OpenConnection(*port)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, pboard)
}
