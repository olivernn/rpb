package main

import (
	"fmt"
	"net"
	"time"
)

const timeout = 2 * time.Second

func OpenConnection(port int) (net.Conn, error) {
	address := fmt.Sprintf("localhost:%d", port)
	connection, err := net.DialTimeout("tcp", address, timeout)

	if err != nil {
		return nil, err
	}

	deadline := time.Now().Add(timeout)
	err = connection.SetDeadline(deadline)

	if err != nil {
		return nil, err
	}

	return connection, nil
}
