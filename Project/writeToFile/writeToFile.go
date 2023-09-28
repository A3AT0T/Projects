package writeToFile

import (
	"Project/model"
	"encoding/json"
	"fmt"

	"log"
	"os"
)

const pass = "D:\\myGO\\Lessons\\Projects\\Project\\storage\\"

func WriteToFileJson(myUser []model.MyUser, filename string) { // write to file format "txt"

	err := os.MkdirAll(pass, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// f, err := os.Create(pass + filename)
	f, err := os.OpenFile(pass+filename+".txt", os.O_APPEND|os.O_CREATE, os.ModeAppend|0777)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	data, err := json.Marshal(myUser)
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = f.Write(data)

	if err != nil {
		log.Fatal(err)
	}
	err = f.Sync()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data written to file successfully")
	fmt.Println(pass + filename)
}

func WriteToFileBinary(encodeusers []byte, filename string) { // write to file format "txt"

	err := os.MkdirAll(pass, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(pass+filename+".txt", os.O_APPEND|os.O_CREATE, os.ModeAppend|0777)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.Write(encodeusers)

	if err != nil {
		log.Fatal(err)
	}
	err = f.Sync()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data written to file successfully")
	fmt.Println(pass + filename + ".txt")
}

func CreateFileName() string {
	var filename string
createFileName:
	fmt.Print("Select filename : \n>")
	fmt.Scanf("%s\n", &filename)
	if filename == "" {
		fmt.Printf("filename can'n be empty !\n")
		goto createFileName
	}

	return filename
}
