package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

// BEGIN1 OMIT
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	herror(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	herror(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	herror(err)
	result, err := ioutil.ReadAll(conn)
	herror(err)
	fmt.Println(string(result))
	os.Exit(0)
}

func herror(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

// END1 OMIT
