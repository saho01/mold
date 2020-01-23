package project

import (
	"bytes"
	"fmt"
	"text/template"
)

func (p *Project) ParseTemplate(f string) error {
	vars := make(map[string]interface{})
	vars["Project"] = fmt.Sprintf("%s/%s", p.repo, p.name)
	vars["Name"] = p.name
	vars["Repo"] = p.repo
	tmpl, err := template.ParseFiles(fmt.Sprintf("templates/%s", f))

	if err != nil {
		return err
		// panic(err)
	}
	// return process(tmpl, vars)
	var tmplBytes bytes.Buffer

	err = tmpl.Execute(&tmplBytes, vars)
	if err != nil {
		return err
		// panic(err)
	}
	// return tmplBytes.String(), nil

	return p.CreateFile(f, tmplBytes.String())

	// func (p *Project) CreateFile(n, content string) error {
}
