package optionWithUsers

import (
	"Project/defaultDatabase"
	"Project/model"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strings"
)

func ConvertUserFieldsWithLimitUser(u []model.User) (mu []model.MyUser, c, index uint8) { // Data standardization
	myUsers := make([]model.MyUser, len(defaultDatabase.Users()))
	var IndexActiveUser uint8
	sourceLenght := 14
	var count uint8
	const LimitUsers = 8
	for i, user := range u {

		if i == LimitUsers {
			break
		}
		count++

		myUsers[i] = model.MyUser{
			Name: user.Name}
		if len(user.Name) > sourceLenght {
			myUsers[i].Name = u[i].Name[:11] + "..."
		}
		if len(user.Name) <= 1 {
			myUsers[i].Name = "NO NAME"
		}

		if !user.Active {
			myUsers[i].Active = false

		}
		if user.Active {
			myUsers[i].Active = true
			IndexActiveUser = (IndexActiveUser | (1 << i))
		}

		switch {
		case user.Age < 0:
			myUsers[i].Age = uint64(user.Age + 100)
		case user.Age > 100:
			myUsers[i].Age = uint64(user.Age - 100)
		case user.Age >= 0:
			myUsers[i].Age = uint64(user.Age)
		}

		switch {
		case user.Mass > 200:
			myUsers[i].Mass = user.Mass * 0.0283495231
		case user.Mass < 1:
			myUsers[i].Mass = user.Mass * 100
		case user.Mass > 1:
			myUsers[i].Mass = user.Mass

		}

		myUsers[i].Books = user.Books

	}

	return myUsers, count, IndexActiveUser
}

func AvgAgeByBook(u []model.MyUser) { // Sorting book by user age
	uniqueBooks := make(map[string]uint64)
	countReader := make(map[string]uint64)

	for _, u := range u {
		for _, book := range u.Books {
			if _, ok := uniqueBooks[book]; !ok {
				uniqueBooks[book] = 0
			}
			uniqueBooks[book] += u.Age
			countReader[book]++

		}
	}
	var lenBookName int
	for v := range uniqueBooks {
		if lenBookName < len(v) {
			lenBookName = len(v)
		}
	}
	titleFormatStatistic := fmt.Sprintf("%%%ds | %%11s\n", lenBookName)
	formatRes := fmt.Sprintf("%%%ds | %%-11d\n", lenBookName)

	fmt.Printf(titleFormatStatistic, "Name book", "Average age")

	for b := range uniqueBooks {
		avg := (uniqueBooks[b]) / (countReader[b])
		fmt.Printf(formatRes, b, avg)
	}
	fmt.Println(strings.Repeat("-", 75))
}

func FindNearbyMass(u []model.MyUser, targetMass float64) float64 { // Finding nearby target by target mass
	nearby := u[0].Mass
	minDiff := math.Abs(float64(targetMass - nearby))

	for i := range u {

		diff := math.Abs(float64(targetMass - u[i].Mass))
		if diff < minDiff {
			nearby = u[i].Mass
			minDiff = diff
		}
	}
	return nearby
}

// func Finder(u []model.MyUser, finder float64) { //Finding and print user with target mass
// 	slices.SortFunc(u, func(a MyUser, b MyUser) bool {
// 		if a.Mass < b.Mass {
// 			return true
// 		}
// 		return false
// 	})
// 	i, _ := slices.BinarySearchFunc(
// 		u,
// 		finder,
// 		func(u MyUser, finder float64) int {
// 			if u.Mass == finder {
// 				return 0
// 			}
// 			if u.Mass > finder {
// 				return 1
// 			}
// 			return -1
// 		},
// 	)
// 	fmt.Printf("Nearby user with mass 80kg is %q he have mass: %.f\n", u[i].Name, u[i].Mass)
// }

func EncodeUsers(mu []model.MyUser) []byte { // Ecoding data to binary format

	buf := bytes.NewBuffer(nil)
	w := func(x any) error {
		w := func(x any) error {
			return binary.Write(buf, binary.BigEndian, x)
		}
		switch x := x.(type) {
		case string:
			err := w(uint8(len(x)))
			if err != nil {
				return err
			}
			err = w([]byte(x))
			if err != nil {
				return err
			}
			return nil
		}
		return w(x)
	}

	for _, u := range mu {
		w(u.Name)
		if u.Active {
			u.Age |= (1 << 63)
		}
		w(u.Age)
		w(u.Mass)
		w(strings.Join(u.Books, ","))
	}
	return buf.Bytes()
}

func DecodeUser(b []byte) ([]model.MyUser, error) { //Decoding data from binary format
	buf := bytes.NewBuffer(b)
	r := io.Reader(buf)
	out := []model.MyUser{}

	for {
		MyUser := model.MyUser{}

		var nameLen uint8
		if err := binary.Read(r, binary.BigEndian, &nameLen); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if nameLen == 0 {
			break
		}

		nameBytes := make([]byte, nameLen)
		if err := binary.Read(buf, binary.BigEndian, &nameBytes); err != nil {
			return nil, err
		}
		MyUser.Name = string(nameBytes)

		// Read Age and Active
		var ageAndActive uint64
		if err := binary.Read(buf, binary.BigEndian, &ageAndActive); err != nil {
			return nil, err
		}
		MyUser.Age = ageAndActive & 0b_0111_1111_1111_1111_1111_1111_1111_1111
		MyUser.Active = (ageAndActive & (1 << 63)) > 0

		// Read Mass
		if err := binary.Read(buf, binary.BigEndian, &MyUser.Mass); err != nil {
			return nil, err
		}

		// Read Books
		var booksLen uint8
		if err := binary.Read(buf, binary.BigEndian, &booksLen); err != nil {
			return nil, err
		}

		if booksLen > 0 {
			booksData := make([]byte, booksLen)
			if err := binary.Read(buf, binary.BigEndian, &booksData); err != nil {
				return nil, err
			}
			MyUser.Books = strings.Split(string(booksData), ",")
		}

		out = append(out, MyUser)
	}

	return out, nil
}

func CountUser(c uint8) (res uint8) { //counting active user
	var countUser uint8
	for i := 0; i < 8; i++ {
		if c&(1<<i) > 0 {
			countUser++
		}
	}
	return countUser

}
