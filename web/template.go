package web
import (
	"html/template"
	"path/filepath"
	"os"
)

func loadTemplate(name string) *template.Template {
	var armagnacPath = os.Getenv("ARMAGNAC_PATH")
	templatePath := filepath.Join(armagnacPath, "templates", name)
	return template.Must(template.ParseFiles(templatePath))
}
