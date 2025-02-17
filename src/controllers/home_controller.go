package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type HomeController struct {
	templates *template.Template
}

func NewHomeController() *HomeController {
	return &HomeController{
		templates: template.Must(template.ParseGlob(filepath.Join("../views", "*.html"))),
	}
}

func (c *HomeController) Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Welcome to Go Web App",
	}
	c.templates.ExecuteTemplate(w, "index.html", data)
}
