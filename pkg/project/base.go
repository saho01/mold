package project

import (
	"go/build"
	"os"
)

type Project struct {
	name   string
	repo   string
	gopath string
}

func NewProject(n, r string) *Project {
	return &Project{
		name:   n,
		repo:   r,
		gopath: getGoPath(),
	}
}

func getGoPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}
