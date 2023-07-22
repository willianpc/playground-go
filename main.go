package main

import (
	"fmt"

	"github.com/willianpc/example/pkg1"
	"github.com/willianpc/example/pkg2"
)

func main() {
	res := pkg1.ApiOne()
	res2 := pkg2.ApiTwo()

	fmt.Println(res)
	fmt.Println(res2)
}
