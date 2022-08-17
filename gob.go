package goblin

import (
	"encoding/gob"
	"os"
)

func GOBWrite(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}

// GOBRead data parameter must be a pointer
func GOBRead(filename string, data interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
