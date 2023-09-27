package serverClient

import (
	"fmt"
	"io/ioutil"

	"net/http"
)

func RunHTTP() {

	// go func() {
	// 	time.Sleep(time.Second)
	// 	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	// 	h := "GET / HTTP/1.1\n"
	// 	h += "Host: 127.0.0.1:8080\n"
	// 	h += "\n"
	// 	conn.Write([]byte(h))
	// 	conn.Close()
	// }()

	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.RequestURI == "/user" {

			body, err := ioutil.ReadFile(`D:\myGO\Lessons\Project\storage\rr.txt`)
			if err != nil {
				fmt.Print(err)
			}

			w.Write([]byte(body))
		}
	},
	),
	)

}
