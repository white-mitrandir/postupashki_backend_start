package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("Это что, брух, что с портом")
		os.Exit(1)
	}
	defer conn.Close()
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("Это что, брух. принять не смогли")
		os.Exit(1)
	}
	if response != "OK\n" {
		fmt.Printf("брух, чет не то, это ваще не окей")
		fmt.Printf(response)
		os.Exit(1)
	} else {
		fmt.Printf("брух, все гуд")
	}
}
