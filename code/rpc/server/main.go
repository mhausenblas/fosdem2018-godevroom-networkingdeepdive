package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

// Arguments represents the input.
type Arguments struct {
	A, B int
}

// Arith represents an arithmetic operation
type Arith int

// Multiply takes two integers and returns the result of their muliplication.
func (t *Arith) Multiply(args *Arguments, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// BEGIN1 OMIT
func main() {
	arith := new(Arith)
	_ = rpc.Register(arith)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// END1 OMIT
