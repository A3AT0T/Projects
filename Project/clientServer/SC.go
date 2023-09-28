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
)

func Sc() {
	ch := make(chan int)
	go func() {
		defer func() {
			ch <- 1
		}()

		listener, err := net.Listen("tcp", "127.0.0.1:8080")
		if err != nil {
			fmt.Println(err)
		}
		defer listener.Close()
		fmt.Println("server.ok ")
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				conn.Close()
				continue
			}

			handlerConn(conn)

		}

	}()

	go func() {
		defer func() {
			ch <- 1
		}()
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			fmt.Println(err)
		}
		defer conn.Close()
		fmt.Println("client.ok")

		var op string

		for {
			fmt.Print("Select operation:\n 1) Add\n 2) Show users\n q) Exit\n> ")

			fmt.Scanln(&op)
			input := make([]byte, 1024)
			input = []byte(op)
			conn.Write(input)

			buff := make([]byte, 1024)
			conn.Read(buff)

		}
	}()
	<-ch
	<-ch
}
func handlerConn(conn net.Conn) {
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
		case "3", "q":
			log.Fatal()
		}
		conn.Write([]byte(op))
	}

}
