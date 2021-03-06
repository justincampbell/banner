package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port = flag.String("port", "8080", "port to listen on")
)

func init() {
	flag.Parse()
}

func main() {
	http.HandleFunc("/", rootHandler)
	fmt.Println("Listening on", *port)
	http.ListenAndServe(":"+*port, nil)
}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Method, request.RequestURI, request.RemoteAddr)
	fmt.Fprint(response, "hi")
}
