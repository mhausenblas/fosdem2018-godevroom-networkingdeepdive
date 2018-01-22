package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

// BEGINC OMIT
func main() {
	if len(os.Args) != 2 {
		log.Fatal("Need a URL to operate on")
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	// dump HTTP header info:
	b, _ := httputil.DumpResponse(resp, false)
	fmt.Print(string(b))
}

// ENDC OMIT
