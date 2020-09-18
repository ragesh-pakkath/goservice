package main

import ( "fmt"
	 "log"
	 "net/http"
	 )

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world")
}

func handlerHi(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hi from GO")
}

func main(){
	http.HandleFunc("/", handler)
	http.HandleFunc("/hi", handlerHi)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
