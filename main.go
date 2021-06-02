package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	//basic http handler
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello GO World!!!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//http.Error(rw,"Oops",http.StatusBadRequest)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Oops"))
			return
		}
		log.Printf("Data %s from request", d)

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye Go World!!!")
		d, _ := ioutil.ReadAll(r.Body)

		fmt.Fprintf(rw, "Goodbye from Go world %s!!!", d)
	})

	//binding address is the localhost on port 9090
	//http.ListenAndServe("127.0.0.1:9090",nil)
	http.ListenAndServe(":9090", nil)
}
