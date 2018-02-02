package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
)

// Arguments represents the input.
type Arguments struct {
	A, B int
}

// BEGIN1 OMIT
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Arguments{17, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)
}

// END1 OMIT
