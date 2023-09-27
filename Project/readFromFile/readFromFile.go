package readFromFile

import (
	"Project/model"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func ReadFromFile() ([]model.MyUser, error) {
	var filePathName string
	fmt.Println("Enter filePathName\n>")
	fmt.Scanf("%s\n", &filePathName)
	file, err := ioutil.ReadFile(filePathName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	buf := bytes.NewBuffer(file)
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
