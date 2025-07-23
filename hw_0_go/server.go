package main

import (
	"fmt"
	"net"
	"os"
)

func sendMes(conn net.Conn) {
	defer conn.Close()
	_, err1 := conn.Write([]byte("OK\n"))
	if err1 != nil {
		fmt.Println("брух, отправить не удалось", err1)
	}
}

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Это что, брух, что с портом")
		os.Exit(1)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Это что, брух, че не подключается")
			os.Exit(1)
		}
		go sendMes(conn)
	}

}
