package printFormat

import (
	"Project/model"
	"fmt"
	"strings"
)

type MyUser struct {
	Name   string
	Age    uint64
	Active bool
	Mass   float64
	Books  []string
}

func PrintFormat(users []model.MyUser) { // Format to output data

	lenName := 0
	lage := 5
	lactive := 10
	lmass := 8
	lenFavoriteBook := 14
	totalLenLine := 0

	for _, u := range users {

		if lenName < len(u.Name) {
			lenName = len(u.Name)
		}
		ll := strings.Join(u.Books, ", ")
		if lenFavoriteBook < len(ll) {
			lenFavoriteBook = len(ll)
		}

	}
	lenName += 2
	lenFavoriteBook += 5
	totalLenLine = lenName + lage + lactive + lmass + lenFavoriteBook

	titleFormat := fmt.Sprintf("┃%%%ds ┃ %%3s ┃ %%6s ┃ %%6s ┃ %%%ds┃\n", lenName, lenFavoriteBook-4)
	fmt.Printf("┏%s%s\n", strings.Repeat("━", totalLenLine), "┓")

	fmt.Printf(titleFormat, "Name", "Age", "Active", "Mass", "Favorite_Books")
	kgFormat := fmt.Sprintf("┃%%%d.%dq ┃ %%3d ┃ %%6t ┃ %%3.f_kg ┃ %%%ds┃\n", lenName, lenName, lenFavoriteBook-4)
	fmt.Printf("┃%s┃\n", strings.Repeat("━", totalLenLine))

	for _, u := range users {
		listBooks := strings.Join(u.Books, ", ")
		if listBooks == "" {
			listBooks = "---"
		}

		fmt.Printf(kgFormat, strings.TrimSpace(u.Name), u.Age, u.Active, u.Mass, listBooks)
		fmt.Printf("┗%s%s\n", strings.Repeat("━", totalLenLine), "┛")

	}
}
