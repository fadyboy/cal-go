package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data any) {
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("Cloning template:%v", err)
		http.Error(w, "there was an error rendering the page", http.StatusInternalServerError)
	}

	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
		},
	)

	w.Header().Set("Content-Type", "text/html;charset=UTF-8")
	err = tpl.Execute(w, data)
	if err != nil {
		log.Printf("Executing template: %v", err)
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl := template.New(patterns[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return `<input type="hidden" />`
			},
		},
	)
	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		htmlTpl: tpl,
	}, nil
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}

	return t
}

// func ParseFS(fs embed.FS, pattern ...string) (Template, error) {
// 	tmpl, err := template.ParseFS(fs, pattern...)
// 	if err != nil {
// 		log.Printf("parsing template: %v", err)
// 		return Template{}, fmt.Errorf("error parsing template: %v", err)
// 	}
//
// 	return Template{htmlTpl: tmpl}, nil
// }
