package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения: ", err)
		os.Exit(1)
	}
	defer conn.Close()

	response, err := io.ReadAll(conn)
	if err != nil {
		fmt.Println("Ошибка чтения: ", err)
		os.Exit(1)
	}
	if string(response) != "OK\n" {
		fmt.Printf("Неверный ответ сервера: %q\n", string(response))
		os.Exit(1)
	}

	fmt.Print(string(response))
}
