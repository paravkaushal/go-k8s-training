package main

import (
	"flag"
	"fmt"
)

func main() {
	stringPtr := flag.String("input", "Hello World!!!", "Welcome string")
	flag.Parse()
	fmt.Println(*stringPtr)
}
