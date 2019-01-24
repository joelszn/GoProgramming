package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://joelszn.github.io")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// make a byte slice with n amount of empty space
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
