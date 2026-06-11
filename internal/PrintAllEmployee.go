package internal

import (
	"fmt"
	"prac/modules"
)

func PrintAllEmployees(employees []modules.Employee) {
	if len(employees) == 0 {
		fmt.Println("Список сотрудников пуст!")
	} else {
		fmt.Println("\n--- Список всех сотрудников ---")
		for i, emp := range employees {
			fmt.Printf("%d. Имя: %s, Возраст: %d, Должность: %s, Зарплата: %.2f\n",
				i+1, emp.Name, emp.Age, emp.Position, emp.Salary)
		}
	}
}
