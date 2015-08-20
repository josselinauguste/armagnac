package web

import (
	"html/template"
	"net/http"
	"path"
	"path/filepath"
	"runtime"

	"github.com/josselinauguste/armagnac/feeds/query"
	"github.com/josselinauguste/magicbus"
)

var digestTemplate = loadTemplate("digest_email.html")

type digestResource struct {
	bus magicbus.Bus
}

func newDigestResource(bus magicbus.Bus) *digestResource {
	return &digestResource{bus}
}

func (resource digestResource) createAndSendDigestHandler(rw http.ResponseWriter, r *http.Request) {
	query := query.NewNewItemsQuery()
	err := resource.bus.Send(query)
	if err != nil {
		//TODO log
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	presenter := newDigestPresenter(*query)
	digestTemplate.Execute(rw, presenter)
}

func loadTemplate(name string) *template.Template {
	_, executedFileName, _, _ := runtime.Caller(1)
	rootPath := path.Dir(executedFileName)
	templatePath := filepath.Join(rootPath, "templates", name)
	return template.Must(template.ParseFiles(templatePath))
}
