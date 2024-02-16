package main

import (
	"fmt"
	"net/http"

	"github.com/arvindeva/touhouapi-cms/internal/data"
	"github.com/arvindeva/touhouapi-cms/internal/validator"
)

func (app *application) createTouhouHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name      string   `json:"name"`
		Species   string   `json:"species"`
		Abilities []string `json:"abilities"`
	}

	err := app.readJSON(w, r, &payload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return

	}

	v := validator.New()

	// name field validation
	v.Check(payload.Name != "", "name", "must be provided")
	v.Check(len(payload.Name) < 500, "name", "must not be more than 500 bytes")

	// species field validation
	v.Check(payload.Species != "", "species", "must be provided")

	// abilities field validation
	v.Check(validator.Unique(payload.Abilities), "abilities", "must not contain duplicate values")
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", payload)
}

func (app *application) showTouhouHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)

		return
	}

	touhou := data.Touhou{
		ID:        id,
		Name:      "Reimu Hakurei",
		Species:   "Human",
		Abilities: []string{"Flying", "Spiritual Power"},
	}

	// Encode the struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, envelope{"touhou": touhou}, nil)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}
}
