package todo

import (
	"fmt"
	"net"
)

var dict = map[string]string{
	"red":    "красный",
	"green":  "зеленый",
	"blue":   "синий",
	"yellow": "желтый",
	"grey":   "сірий",
}

func ServerClient() {
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
			handleConnection(conn) // запускаем горутину для обработки запроса
		}
	}()
	// обработка подключения

	go func() {
		defer func() {
			ch <- 0
		}()
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()

		for {
			var op string
			fmt.Print("Select operation:\n 1) Add\n 2) Delete\n 3) Search user\n 4) Show all\n q) Exit\n> ")

			fmt.Print("Введите слово: ")
			_, err := fmt.Scanf("%s\n", &op)
			if err != nil {
				fmt.Println("Invalid input", err)
				continue
			}
			// отправляем сообщение серверу
			if n, err := conn.Write([]byte(op)); n == 0 || err != nil {
				fmt.Println(err)
				return
			}
			// получем ответ
			fmt.Print("Перевод:")
			buff := make([]byte, 1024)
			n, err := conn.Read(buff)
			if err != nil {
				break
			}
			fmt.Print(string(buff[0:n]))
			fmt.Println()
		}
	}()

	<-ch
	<-ch

}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		// считываем полученные в запросе данные
		input := make([]byte, 1024)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Read error:", err)
			break
		}
		op := string(input[0:n])
		// на основании полученных данных получаем из словаря перевод
		target, ok := dict[op]

		if ok == false { // если данные не найдены в словаре
			target = "undefined"
		}
		// выводим на консоль сервера диагностическую информацию
		fmt.Println(op, "-", target)
		// отправляем данные клиенту

		conn.Write([]byte(target))
	}
}
