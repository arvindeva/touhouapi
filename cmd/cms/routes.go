package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/touhous", app.createTouhouHandler)
	router.HandlerFunc(http.MethodGet, "/touhous/:id", app.showTouhouHandler)

	// Return the httprouter instance.
	return app.recoverPanic(router)
}