package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
)

func main() {
	port := "8080"

	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer listener.Close()

	localIp, errIP := getLocalIP()
	if errIP != nil {
		fmt.Fprintln(os.Stderr, "An unexpected error has occured retrieving local IP:", errIP.Error())
		os.Exit(1)
	}

	fmt.Printf("Chat server started on %s. Listening on %s\n", localIp, port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		go handleClient(conn)
	}
}

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	return strings.Split(conn.LocalAddr().String(), ":")[0], nil
}
