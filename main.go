package main

import "net/http"
import "fmt"

func HelloHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello word!"))
}

func main() {
	fmt.Println("start server")
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8888", nil)
}

