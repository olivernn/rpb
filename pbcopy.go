package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	port := flag.Int("port", PbcopyPort, "remote pbcopy port")
	showVersion := flag.Bool("v", false, "display version")

	flag.Parse()

	if *showVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	pboard, err := OpenConnection(*port)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(pboard, os.Stdin)
}
