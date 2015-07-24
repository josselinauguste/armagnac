package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/josselinauguste/armagnac/web"
)

func main() {
	r := web.RegisterRoutes()
	http.ListenAndServe(fmt.Sprintf(":%v", getPort()), r)
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return port
}
