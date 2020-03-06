package main

import (
	"fmt"
	"os"

	ecom "github.com/cardenasrjl/ecom/pkg/cmd/server"
)

func main() {
	if err := ecom.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}