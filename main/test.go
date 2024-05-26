package main

import (
	"fmt"
	_ "sync" // Now the error of package imported but not used will not come becasue of "_"
)

func main() {

	hello := "Hello World"
	fmt.Println(hello)
}
