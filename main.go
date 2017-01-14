package main

import (
	"log"
	"net/http"

	"github.com/codegansta/negroni"
	"github.com/udemy/common"
	"github.com/udemy/routers"
)

func main() {
	common.StartUp()
	r := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(r)

	log.Println("Listening...")
	http.ListenAndServe(":8080", n)
}
