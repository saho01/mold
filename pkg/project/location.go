package project

import (
	"fmt"
	"os"
)

const (
	FILE = 0664
	DIR  = 0775
)

func CreateDir(d string) error {
	_, err := os.Stat(d)
	if os.IsNotExist(err) {
		fmt.Printf("creating project at %s\n", d)
		return os.Mkdir(d, os.FileMode(DIR))
	}
	return fmt.Errorf("directory already exists")
}

func CreateFile(d, n string) error {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", d, n), os.O_RDONLY|os.O_CREATE, os.FileMode(FILE))
	if err != nil {
		return err
	}
	return f.Close()
}
