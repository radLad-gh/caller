package main

import (
	"fmt"

	sample "github.com/radLad-gh/sample/src/util"
)

func main() {
	fmt.Println(sample.About())
	fmt.Println(sample.Hello("Justice"))
	fmt.Println(sample.GetQuote())
}
