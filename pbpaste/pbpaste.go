package main

import (
	"flag"
	"fmt"
	"github.com/olivernn/rpb"
	"io"
	"os"
)

func main() {
	port := flag.Int("port", 2225, "remote pbpaste port")
	flag.Parse()
	pboard, err := rpb.OpenRead(*port)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, pboard)
}
