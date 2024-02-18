package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/touhous", app.listTouhousHandler)
	router.HandlerFunc(http.MethodPost, "/touhous", app.createTouhouHandler)
	router.HandlerFunc(http.MethodGet, "/touhous/:id", app.showTouhouHandler)
	router.HandlerFunc(http.MethodPatch, "/touhous/:id", app.updateTouhouHandler)
	router.HandlerFunc(http.MethodDelete, "/touhous/:id", app.deleteTouhouHandler)

	return app.recoverPanic(app.rateLimit(router))
}
