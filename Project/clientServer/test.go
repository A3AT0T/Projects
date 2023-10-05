package serverClient

import (
	"fmt"
	"net"
	"net/http"
)

func TEST() {
	// Start a TCP server
	go startTCPServerr()

	// Register an HTTP handler for a specific endpoint
	http.HandleFunc("/send-to-tcp", func(w http.ResponseWriter, r *http.Request) {
		// Connect to the TCP server

		conn, err := net.Dial("tcp", "localhost:12345") // Replace with your TCP server address
		if err != nil {
			http.Error(w, "Failed to connect to TCP server", http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		// Prepare and send data to the TCP server
		data := []byte("Hello, TCP Server!")
		_, err = conn.Write(data)
		if err != nil {
			http.Error(w, "Failed to send data to TCP server", http.StatusInternalServerError)
			return
		}

		// fmt.Fprintln(w, "Data sent to TCP server successfully")
		fmt.Println("Data sent to TCP server successfully")

	})

	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func startTCPServerr() {
	// Start a simple TCP server on localhost:12345
	ln, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	fmt.Println("TCP server started")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		// Handle incoming data from the TCP client
		go handleTCPClient(conn)
	}
}

func handleTCPClientt(conn net.Conn) {
	// Read data from the TCP client
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Process and log the received data
	receivedData := buffer[:n]
	fmt.Printf("Received from TCP client: %s\n", string(receivedData))
}
