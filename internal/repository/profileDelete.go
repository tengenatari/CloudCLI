package repository

import (
	"fmt"
	"os"
)

func (y *YamlRepository) ProfileDelete(name string) error {
	err := os.Remove(fmt.Sprintf("%s.yaml", name))
	if err != nil {
		return err
	}
	return nil
}
