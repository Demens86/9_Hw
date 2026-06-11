package main

import (
	"fmt"
	"prac/internal"
	"prac/modules"
)

func main() {
	internal.PrintAllEmployees(modules.Employees)
	fmt.Println("---------------")
	err := internal.LoadFile("./test.json")
	if err != nil {
		return
	}
	internal.PrintAllEmployees(modules.Employees)

}
