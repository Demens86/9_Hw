package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"prac/modules"
)

func WriteInFile(employees []modules.Employee) error {
	jsonData, err := json.Marshal(employees)
	if err != nil {
		return err
	}

	err1 := os.WriteFile("./employees.json", jsonData, 0644)
	if err1 != nil {
		return err1
	}
	fmt.Println("Запись в файл выполнена!")
	return nil
}
