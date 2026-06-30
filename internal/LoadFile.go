package internal

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"prac/modules"
)

func LoadFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	//err1 := json.Unmarshal(data, &modules.Employees)
	//if err1 != nil {
	//	return err1
	//}
	dec := json.NewDecoder(bytes.NewReader(data))
	for {
		if err := dec.Decode(&modules.Employees); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

	}

	return nil
}
