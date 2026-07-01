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
		for _, emp := range employees {
			fmt.Printf("%d. Имя: %s, Возраст: %d, Должность: %s, Зарплата: %.2f, Электронная почта: %s, "+
				"Статус: %t\n",
				emp.ID, emp.Name, emp.Age, emp.Position, emp.Salary, emp.Email, emp.IsActive)
		}
	}
}
