package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	_, err := conn.Write([]byte("OK\n"))
	if err != nil {
		fmt.Println("Send error: ", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка подключения к серверу: ", err)
		return
	}
	defer listener.Close()

	fmt.Println("Сервер подключен к порту 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(("Ошибка подключения: "), err)
			continue
		}

		go handleConnection(conn)
	}
}
