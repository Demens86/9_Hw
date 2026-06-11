package internal

import (
	"encoding/json"
	"os"
	"prac/modules"
)

func LoadFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err1 := json.Unmarshal(data, &modules.Employees)
	if err1 != nil {
		return err1
	}

	return nil
}
