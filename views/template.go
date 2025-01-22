package views

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func Parse(templateFilepath string) (Template, error) {
	htmlTpl, err := template.ParseFiles(templateFilepath)
	if err != nil {
		log.Printf("Parsing template: %v", err)
		return Template{}, fmt.Errorf("Error parsing template %v", err)
	}

	return Template{htmlTpl: htmlTpl}, nil
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return t
}

func ParseFS(fs embed.FS, pattern ...string) (Template, error) {
	tmpl, err := template.ParseFS(fs, pattern...)
	if err != nil {
		log.Printf("parsing template: %v", err)
		return Template{}, fmt.Errorf("Error parsing template: %v", err)
	}

	return Template{htmlTpl: tmpl}, nil
}
