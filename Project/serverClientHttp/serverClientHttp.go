package serverClientHttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func RunServerHTTP() {
	http.ListenAndServe(
		":8080",
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				b, _ := ioutil.ReadAll(r.Body)

				fmt.Println(r.Method)

				for k, v := range r.Header {
					fmt.Printf("%s: %s\n", k, v)
				}

				fmt.Println(string(b))

				w.Write([]byte(`hi`))
			},
		),
	)

}

func RunClientHttp() {
	fmt.Println("asd")
}
