package internal

import (
	"math"
	"prac/modules"
)

func ChangeSalary() {
	for i := range modules.Employees {
		temp := modules.Employees[i].Salary * 1.1
		modules.Employees[i].Salary = math.Round(temp)
	}
}
