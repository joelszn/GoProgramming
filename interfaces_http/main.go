package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://joelszn.github.io")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// same thing as prev code just simplified returns HTML
	io.Copy(os.Stdout, resp.Body)
}
