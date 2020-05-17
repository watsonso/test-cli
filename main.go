package main

import (
	"flag"
	"fmt"
)

func main() {
	var name = flag.String("name", "Aoki", "please specify -name flag")
	flag.Parse()
	fmt.Println("This name flag is", *name)
}
