// BEGINC OMIT
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fqdn := os.Args[1]
	a, err := net.LookupHost(fqdn)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}

// ENDC OMIT
