package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type UserController struct {
	templates *template.Template
}

func NewUserController() *UserController {
	return &UserController{
		templates: template.Must(template.ParseGlob(filepath.Join("views", "*.html"))),
	}
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Welcome to Go Web App",
	}
	c.templates.ExecuteTemplate(w, "index.html", data)
}
