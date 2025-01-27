package handlers

import "text/template"

// var templates = template.Must(template.New("").ParseFiles(
//
//	"templates/login.html",
//	"templates/signup.html",
//	"templates/error.html",
//	// "templates/signup.html",
//	// "templates/signup.html",
//	// "templates/signup.html",
//	// "templates/signup.html",
//	// "templates/signup.html",
//
// ))

var templates *template.Template

// Initialize templates (call this function in your main.go or init function)
func InitTemplates(templateDir string) {
	templates = template.Must(template.ParseGlob(templateDir + "/*.html"))
}
