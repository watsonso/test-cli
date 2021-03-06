package main

import (
	"flag"
	"fmt"
)

func main() {
	// flagパッケージについて ref: https://golang.org/pkg/flag
	// flag.String("引数の名前", "デフォルト値", "ヘルプメッセージ")
	var name = flag.String("name", "Aoki", "please specify -name flag")
	flag.Parse()

	// fmtパッケージについて ref: https://golang.org/pkg/fmt/
	fmt.Println("This name flag is", *name)
}
