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
	// dump response body
	var buf [512]byte
	reader := resp.Body
	payload := 0
	for {
		n, err := reader.Read(buf[0:])
		if err != nil {
			break
		}
		payload += n
		fmt.Print(string(buf[0:n]))
	}
	fmt.Printf("\n--DONE reading %d bytes\n", payload)
}

// ENDC OMIT
