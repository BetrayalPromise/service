package main

import (
	"controller"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user/getInfo", controller.HandleLogin)
	http.HandleFunc("/test/api", controller.API)
	var error = http.ListenAndServe("127.0.0.1:8081", nil)
	if error != nil {
		fmt.Println("Service Error: ", error)
	}
}
