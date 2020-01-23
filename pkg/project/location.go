package project

import (
	"fmt"
	"os"
)

const (
	FILE = 0664
	DIR  = 0775
)

func (p *Project) CreateDir() error {
	path := fmt.Sprintf("%s/src/%s/%s", p.gopath, p.repo, p.name)
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		fmt.Printf("creating project at %s\n", path)
		return os.MkdirAll(path, os.FileMode(DIR))
	}
	return fmt.Errorf("directory already exists")
}

func (p *Project) CreateFile(n, content string) error {
	path := fmt.Sprintf("%s/src/%s/%s", p.gopath, p.repo, p.name)
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", path, n), os.O_RDWR|os.O_CREATE, os.FileMode(FILE))
	if err != nil {
		return err
	}
	if content != "" {
		_, err := f.WriteString(content)
		if err != nil {
			return err
		}
	}
	return f.Close()
}
