package web

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/josselinauguste/armagnac/feeds/query"
	"github.com/josselinauguste/magicbus"
)

type digestResource struct {
	bus    magicbus.Bus
	mailer mailer
}

var digestTemplate = loadTemplate("digest_email.html")

func newDigestResource(bus magicbus.Bus, mailer mailer) *digestResource {
	return &digestResource{bus, mailer}
}

func (resource digestResource) createAndSendDigestHandler(rw http.ResponseWriter, r *http.Request) {
	digest, err := resource.createDigest()
	if err != nil {
		fmt.Printf("ERROR: %#v\n", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	err = resource.mailer.sendMail("jauguste@iblop.net", "A week digested", digest)
	if err != nil {
		fmt.Printf("ERROR: error while sending digest by email: %#v\n", err.Error())
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}

func (resource digestResource) createDigest() ([]byte, error) {
	query := query.NewNewItemsQuery()
	err := resource.bus.Send(query)
	if err != nil {
		return nil, fmt.Errorf("error while sending to bus: %#v\n", err.Error())
	}
	presenter := newDigestPresenter(*query)
	buffer := new(bytes.Buffer)
	err = digestTemplate.Execute(buffer, presenter)
	if err != nil {
		return nil, fmt.Errorf("error while generating reponse body: %#v\n", err.Error())
	}
	return buffer.Bytes(), nil
}
