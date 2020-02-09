package main

import (
	"fmt"
	"os"

	"github.com/cardenasrjl/ecom/pkg/cmd"
)

func main() {
	if err := RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}