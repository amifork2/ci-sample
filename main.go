package main

import (
	"fmt"

	"github.com/amifork2/sample/add"
	"github.com/amifork2/sample/sub"
)

func main() {
	fmt.Println(add.Add(10, 20))
	fmt.Println(sub.Sub(10, 20))
}
