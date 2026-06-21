// парсинг json внесение изменений и запись в другой json файл
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

	fmt.Println()
	fmt.Println("---------------")

	internal.ChangeSalary()
	err1 := internal.WriteInFile(modules.Employees)
	if err1 != nil {
		return
	}
	fmt.Println("---------------")
	internal.PrintAllEmployees(modules.Employees)

}
