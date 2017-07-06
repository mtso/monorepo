package main

import (
	"fmt"

	"github.com/dop251/goja"
)

func throw() {
	vm := goja.New()
	_, err := vm.RunString("2+2; throw new Error('test')")
	panic(err)
}

func main() {
	vm := goja.New()
	v, err := vm.RunString("2+2")
	if err != nil {
		panic(err)
	}
	if num := v.Export().(int64); num != 4 {
		panic(num)
	} else {
		fmt.Println(num)
	}
}
