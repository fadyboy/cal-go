package controllers

import (
	"html/template"
	"net/http"

	"github.com/fadyboy/lenslocked/views"
)

type Static struct {
	Template views.Template
}

func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tmpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes, there is a free trial for 30 days",
		},
		{
			Question: "What are your support hours?",
			Answer:   "9AM - 5PM EST, Monday - Friday",
		},
		{
			Question: "How do I contact support?",
			Answer:   `You can send an email to <a href="mailto:support@example.com">support@example.com</a>`,
		},
		{
			Question: "What is your refund policy?",
			Answer:   "You can get a refund after 14 days",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, questions)
	}
}
