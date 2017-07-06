package main

import (
	"fmt"
	"time"

	"github.com/dop251/goja"
)

// Measure time from "goja.New()" to when RunString()
// or RunProgram() returns.
func main() {
	runstring()
	runprogram()
}

func runstring() {
	start := time.Now()

	vm := goja.New()
	_, err := vm.RunString("2+2")
	if err != nil {
		panic(err)
	}

	diff := time.Now().Sub(start)
	fmt.Println("vm.RunString():", diff)
}

func runprogram() {
	prog, err := goja.Compile("two plus two", "2+2", false)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	vm := goja.New()
	_, err = vm.RunProgram(prog)
	if err != nil {
		panic(err)
	}
	diff := time.Now().Sub(start)
	fmt.Println("vm.RunProgram():", diff)
}
