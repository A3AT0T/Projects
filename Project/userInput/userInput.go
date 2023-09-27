package userInput

import (
	"Project/model"
	"Project/printFormat"

	"fmt"
)

var filename string

func UserInput() []model.MyUser { // simple input without recursion & create file

	var op string
	var name string
	var age uint64
	var active bool
	var mass float64
	var book string
	var slbooks []string

	users := []model.MyUser{}

	for {

		fmt.Print("Select operation:\n 1) Add\n 2) Delete\n 3) Search user\n 4) Show all\n q) Exit\n> ")
		fmt.Scanf("%s\n", &op)

		switch op {
		case "1", "Add", "add":
			fmt.Print("Enter User Name\n>  ")
		startToNameCheck:
			fmt.Scanf("%s\n", &name)
			for _, user := range users {
				if user.Name == name {

					fmt.Print("Sorry, but user with this name was created before, enter other name\n> ")
					goto startToNameCheck

				}
				break

			}

			fmt.Print("Enter Age\n>  ")
			fmt.Scanf("%d\n", &age)
			fmt.Print("Enter Active\n>  ")
			fmt.Scanf("%t\n", &active)
			fmt.Print("Enter Mass\n>  ")
			fmt.Scanf("%f\n", &mass)
			fmt.Print("Enter Favorite_Books\n  ")

			for {
				fmt.Print("Enter books name(for exit enter 'q')\n> ")
				fmt.Scanf("%s\n", &book)

				if book == "q" || book == "" {
					goto endScanBook
				}
				slbooks = append(slbooks, book)

			}
		endScanBook:
			users = append(users, model.MyUser{Name: name, Age: age, Active: active, Mass: mass, Books: slbooks})
			slbooks = []string{}
			age = 0
			mass = 0

		case "2", "Delete", "delete", "del":
			var denNlname string
			var detectorUser int
			fmt.Print("Enter name\n>  ")
			fmt.Scanf("%s\n", &denNlname)
			var indexUser int
			for i, user := range users {
				if user.Name == denNlname {
					indexUser = i
					users = append(users[:indexUser], users[indexUser+1:]...)
					fmt.Println("User deleted successfully.")
					detectorUser++
					break
				}

			}
			if detectorUser == 0 {
				fmt.Println("User not found!!!\n")
			}

		case "3", "find", "search":
			fmt.Scanf("%s\n", &name)
			fmt.Print("Enter user name\n>  ")
			for i, user := range users {
				if user.Name == name {
					fmt.Printf("i:=%d\n name: %s\n", i, user.Name)
					break
				}
				fmt.Println("no name")
				break
			}

		case "4", "show", "sh":
			printFormat.PrintFormat(users)

		case "q", "Exit", "exit", "quit":
			goto exit
		}
	}
exit:
	return users
}
