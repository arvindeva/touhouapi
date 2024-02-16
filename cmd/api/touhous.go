package main

import (
	"fmt"
	"net/http"

	"github.com/arvindeva/touhouapi/internal/data"
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
