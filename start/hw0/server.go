package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func run() error {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	defer listener.Close()

	log.Println("Server is listening on port 8080")

	var tempDelay time.Duration

	for {
		conn, err := listener.Accept()

		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}

			if tempErr, ok := err.(net.Error); ok && tempErr.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}

				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}

				log.Printf(
					"Temporary accept error: %v; retrying in %v",
					err,
					tempDelay,
				)

				time.Sleep(tempDelay)
				continue
			}

			return fmt.Errorf("accept: %w", err)
		}

		tempDelay = 0

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	if _, err := conn.Write([]byte("OK\n")); err != nil {
		log.Printf("Write error: %v", err)
	}
}
