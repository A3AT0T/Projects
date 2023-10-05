package serverClient

import (
	// "fmt"
	// "io/ioutil"

	"fmt"

	"net"

	"net/http"
)

func RunHTTP() {
	go startTCPServer()
	// check uri & go
	http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI == "/user" {
			conn, err := net.Dial("tcp", "127.0.0.1:12345")
			if err != nil {
				http.Error(w, "Failed to connect TCP server", http.StatusInternalServerError)
				return
			}
			defer conn.Close()
			// Prepare and send data to the TCP server
			data := []byte("show me my user")
			_, err = conn.Write(data)
			if err != nil {
				http.Error(w, "Failed to send data to TCP srerve", http.StatusInternalServerError)
			}
			return
		}
		fmt.Fprintln(w, "Data sent to TCP server successfully")

	}))
}

func startTCPServer() {
	// Start a simple TCP server on localhost:12345
	l, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Println("TCP server started")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		// Handle incoming data from the TCP client
		go handleTCPClient(conn)
	}
}

func handleTCPClient(conn net.Conn) {
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
	if string(receivedData) == "show me my user" {

	}
}
