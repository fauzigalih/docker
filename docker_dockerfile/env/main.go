package main

import {
	"fmt"
	"net/http"
	"os"
}

func main() {
	port := os.Getenv("APP_PORT")
	fmt.Println("Run app in : "+port)
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":"+port,nil)
}

func HelloServer(w http.ResponsWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}