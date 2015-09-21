package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func OpenRead(port int) (io.Reader, error) {
	address := fmt.Sprintf("localhost:%d", port)
	r, err := net.DialTimeout("tcp", address, 2*time.Second)

	if err != nil {
		return nil, err
	}

	err = r.SetReadDeadline(deadline())

	if err != nil {
		return nil, err
	}

	return r, nil
}

func OpenWrite(port int) (io.Writer, error) {
	address := fmt.Sprintf("localhost:%d", port)
	w, err := net.DialTimeout("tcp", address, 2*time.Second)

	if err != nil {
		return nil, err
	}

	err = w.SetWriteDeadline(deadline())

	if err != nil {
		return nil, err
	}

	return w, nil
}

func deadline() time.Time {
	return time.Now().Add(2 * time.Second)
}
