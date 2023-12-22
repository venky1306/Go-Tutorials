package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Listen on TCP port 8080 on all interfaces.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error Creating tcp listener on port 8080.")
	}
	defer listener.Close()

	fmt.Println("Listening TCP on :8080.")

	// Listen on UDP port 8080 on all interfaces.
	udpserver, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		log.Fatal("Error Listening on UDP.")
	}
	defer udpserver.Close()

	fmt.Println("Listening UDP on :8080")

	go func() {
		for {
			// Accept a TCP connection.
			connection, err := listener.Accept()
			if err != nil {
				log.Panic("Unable to accept connection.")
				continue
			}
			go handleConn(connection)
		}
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			// Wait for a connection and read data from the client.
			n, addr, err := udpserver.ReadFrom(buf)
			if err != nil {
				fmt.Println("Error reading data", err.Error())
				continue
			}
			go handlePacket(udpserver, addr, buf[:n])
		}
	}()

	// Block forever to keep the main thread from exiting.
	select {}
}

func handleConn(connection net.Conn) {
	defer connection.Close()

	buf := make([]byte, 1024)
	// Receive data from the client.
	x, err := connection.Read(buf)
	if err != nil {
		fmt.Println("Unable to read incomming data.", err.Error())
	}

	fmt.Printf("value of x %d, and received data: %s \n", x, string(buf))

	// Send data to the client.
	_, err = connection.Write(buf)
	if err != nil {
		fmt.Println("Error sending data through TCP.", string(buf))
	}

}

func handlePacket(udpConn net.PacketConn, addr net.Addr, p []byte) {
	fmt.Printf("Received data: %s \n", string(p))

	// Send data to the client.
	_, err := udpConn.WriteTo(p, addr)
	if err != nil {
		fmt.Println("Error sending data through UDP packet.", string(p))
	}
}
