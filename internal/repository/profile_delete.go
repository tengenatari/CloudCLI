package repository

import (
	"fmt"
	"os"
)

func (y *YamlRepository) ProfileDelete(name string) error {
	profile, _ := LoadProfile(name, y.dir)
	if profile == nil {
		return fmt.Errorf("profile %s not found", name)
	}

	err := os.Remove(fmt.Sprintf("%s/%s.yaml", y.dir, name))
	if err != nil {
		return err
	}
	return nil
}
