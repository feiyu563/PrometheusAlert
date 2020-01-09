package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	if len(os.Args) < 2 {
		return
	}
	switch os.Args[1] {
	case "endpoint":
		err = endpointHandle(os.Args[1:])
	}
	if err != nil {
		fmt.Println(err)
	}
}
