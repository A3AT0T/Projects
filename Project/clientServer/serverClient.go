package serverClient

import (
	"Project/model"
	"Project/optionWithUsers"
	"Project/printFormat"
	"Project/readFromFile"
	"Project/userInput"
	"Project/writeToFile"
	"fmt"
	"net"
)

func ServerClient() {
	//server side

	ch := make(chan int)

	go func() {
		defer func() {
			ch <- 0
		}()
		listener, err := net.Listen("tcp", ":8080")

		if err != nil {
			fmt.Println(err)
			return
		}
		defer listener.Close()
		fmt.Println("Server is listening...")
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println(err)
				conn.Close()
				continue
			}

			handleConnection(conn)
		}

	}()

	// client side
	go func() {
		defer func() {
			ch <- 1
		}()
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("The connection is established")
		defer conn.Close()

		for {
			var op string
			fmt.Print("Select operation:\n 1) Add\n 2) Show users\n q) Exit\n> ")

			_, err := fmt.Scanf("%s\n", &op)
			if err != nil {
				fmt.Println("Invalid input", err)
				continue
			}
			if n, err := conn.Write([]byte(op)); n == 0 || err != nil {
				fmt.Println(err)
				return
			}

			buff := make([]byte, 1)
			n, err := conn.Read(buff)
			if err != nil {
				break
			}
			_ = n

		}
	}()

	// go func() {
	// 	defer func() {
	// 		ch <- 0
	// 	}()
	// 	RunHTTP()
	// }()

	<-ch
	<-ch

}

// handler

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, 1024)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		op := string(input[0:n])

		switch op {
		case "1", "start":
			var MU []model.MyUser
			MU = userInput.UserInput()

			fmt.Print("Adding users is done, if you want write to file enter: \"save\":\n")
			var save string
			fmt.Scanf("%s", &save)

			if save != "save" {
				fmt.Print("file wasn't created\n")
				break
			}

			filename := writeToFile.CreateFileName()
			writeToFile.WriteToFileBinary(optionWithUsers.EncodeUsers(MU), filename)
			fmt.Print("Successfully\n")
		case "2", "show":
			mu, err := readFromFile.ReadFromFile()
			if err != nil {
				fmt.Println(err)
			}
			printFormat.PrintFormat(mu)
		case "q", "exit":
			break //todo
		}
		conn.Write([]byte(op))

	}
}
