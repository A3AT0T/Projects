package serverClient

import (
	"Project/optionWithUsers"
	"Project/printFormat"
	"Project/readFromFile"
	"Project/userInput"
	"Project/writeToFile"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
)

func handleClient(conn net.Conn) {
	defer conn.Close()

	for {

		buf := make([]byte, 1024)
		n, err := conn.Read([]byte(buf))
		if err != nil {
			fmt.Println(err)
		}
		op := (string(buf[0:n]))
		switch op {
		case "1", "add":
			fmt.Println("command 1")
			mu := userInput.UserInput()
			fmt.Print("Adding users is done, if you want write to file enter: \"save\":\n")
			var save string
			fmt.Scanf("%s", &save)

			if save != "save" {
				fmt.Print("file wasn't created\n")
				break
			}

			filename := writeToFile.CreateFileName()
			writeToFile.WriteToFileBinary(optionWithUsers.EncodeUsers(mu), filename)
			fmt.Print("Successfully\n")

		case "2", "show":
			mu, err := readFromFile.ReadFromFile()
			if err != nil {
				fmt.Println(err)
			}
			printFormat.PrintFormat(mu)

		case "111":
			fmt.Println("when pigs fly")

		case "3", "q":
			log.Fatal()
		}
		conn.Write([]byte(op))
	}

	// Start an HTTP server on the client side within a Goroutine
	go func() {
		http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello from the client's HTTP server!")
		})
		fmt.Println("Client HTTP server listening on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("client.ok")

	var op string

	for {
		fmt.Print("Select operation:\n 1) Add\n 2) Show users\n q) Exit\n> ")

		_, err := fmt.Scanln(&op)
		if err != nil {
			fmt.Println("Invalid input", err)
			continue
		}
		input := make([]byte, 1024)
		input = []byte(op)
		conn.Write(input)

		buff := make([]byte, 1024)
		conn.Read(buff)

	}
}

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server is listening on :12345")

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			handleClient(conn)
		}()
	}

	wg.Wait()
}
