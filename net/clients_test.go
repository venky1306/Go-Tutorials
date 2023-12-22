package main

import (
	"fmt"
	"net"
	"testing"
)

func TestTcp(t *testing.T) {

	for i := 0; i < 100; i++ {
		// Connect to the remote TCP server.
		conn, err := net.Dial("tcp", ":8080")
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		// Close the connection when done.
		defer conn.Close()
		fmt.Println("Connected to", conn.RemoteAddr())

		// Send data to the server.
		_, err = conn.Write([]byte(fmt.Sprintf("Test #%d", i)))
		if err != nil {
			fmt.Println("Error sending:", err.Error())
			return
		}

		// Receive data from the server.
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("Error receiving:", err.Error())
			return
		}
		fmt.Println("Received data:", string(buf))
	}
}

func TestUdp(t *testing.T) {

	for i := 0; i < 100; i++ {
		// Connect to the remote UDP server.
		conn, err := net.Dial("udp", ":8080")
		if err != nil {
			fmt.Println("Error Dailing:", err.Error())
		}
		defer conn.Close()

		fmt.Println("Connected", conn.RemoteAddr())

		// Send data to the server.
		_, err = conn.Write([]byte(fmt.Sprintf("Test #%d", i)))
		if err != nil {
			fmt.Println("Error sending:", err.Error())
			return
		}

		buf := make([]byte, 1024)
		// Receive data from the server.
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("Error receiving:", err.Error())
			return
		}
		fmt.Println("Received data:", string(buf))
	}
}
