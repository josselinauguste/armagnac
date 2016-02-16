package web

import (
	"html/template"
	"os"
	"path/filepath"
)

func loadTemplate(name string) *template.Template {
	var armagnacPath = os.Getenv("ARMAGNAC_PATH")
	templatePath := filepath.Join(armagnacPath, "templates", name)
	return template.Must(template.ParseFiles(templatePath))
}
