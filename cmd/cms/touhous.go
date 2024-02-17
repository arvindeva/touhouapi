package main

import (
	"errors"
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

	touhou := &data.Touhou{
		Name:      payload.Name,
		Species:   payload.Species,
		Abilities: payload.Abilities,
	}

	v := validator.New()

	if data.ValidateTouhou(v, touhou); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Touhous.Insert(touhou)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/touhous/%d", touhou.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"touhou": touhou}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) showTouhouHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	touhou, err := app.models.Touhous.Get(id)

	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Encode the struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK, envelope{"touhou": touhou}, nil)
}
