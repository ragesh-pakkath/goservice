package main

import ( "fmt"
	 "log"
	 "net/http"
	 )

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello world")
}

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
