package handler

import (
	"net/http"
	"regexp"
	"text/template"
)

var html *template.Template
var text *template.Template

func init() {
	html = template.Must(template.ParseGlob("templates/html/*.gohtml"))
	text = template.Must(template.ParseGlob("templates/text/*.txt"))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	rgx := regexp.MustCompile("^curl")

	if rgx.MatchString(r.UserAgent()) {
		text.ExecuteTemplate(w, "test.txt", nil)
	} else {
		html.ExecuteTemplate(w, "index.gohtml", nil)
	}
}
