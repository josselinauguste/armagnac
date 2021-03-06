package web

import (
	"github.com/gorilla/mux"
	"github.com/josselinauguste/magicbus"
)

func RegisterRoutes() *mux.Router {
	bus := getBus()
	mailer := htmlMailer{}
	digestResource := newDigestResource(bus, mailer)
	r := mux.NewRouter().StrictSlash(false)
	digests := r.Path("/digests").Subrouter()
	digests.Methods("POST").HandlerFunc(digestResource.createAndSendDigestHandler)
	return r
}

func getBus() magicbus.Bus {
	return magicbus.NewSynchronousBus()
}
