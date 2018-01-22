package main

// BEGIN1 OMIT
import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Need a URL and a limit (in bytes) to operate on.")
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err.Error())
	}
	// dump HTTP header info:
	b, _ := httputil.DumpResponse(resp, false)
	fmt.Print(string(b))
	// END1 OMIT
	// BEGIN2 OMIT
	// dump response body
	var buf [512]byte
	reader := resp.Body
	payload := 0
	limit, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	for payload < limit {
		n, err := reader.Read(buf[0:])
		if err != nil {
			break
		}
		payload += n
		fmt.Print(string(buf[0:n]))
	}
	fmt.Printf("\n--DONE reading %d bytes\n", payload)
	// END2 OMIT
}
