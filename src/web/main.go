package main

import (
	"fmt"
	"net/http"
)

/**
 * @author: DJ
 * @date: 2024-12-13  21:47
 * @description:
 */

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello DJ")
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
		return
	}
}
