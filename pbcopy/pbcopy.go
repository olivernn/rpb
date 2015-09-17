package main

import (
	"flag"
	"fmt"
	"github.com/olivernn/rpb"
	"io"
	"os"
)

func main() {
	port := flag.Int("port", 2224, "remote pbcopy port")
	flag.Parse()
	pboard, err := rpb.OpenWrite(*port)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(pboard, os.Stdin)
}
