package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/arvindeva/touhouapi-cms/internal/data"
	"github.com/arvindeva/touhouapi-cms/internal/validator"
)

func (app *application) listTouhousHandler(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Name    string
		Species string
		data.Filters
	}

	v := validator.New()

	qs := r.URL.Query()

	payload.Name = app.readString(qs, "name", "")
	payload.Species = app.readString(qs, "species", "")

	payload.Filters.Page = app.readInt(qs, "page", 1, v)
	payload.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	payload.Filters.Sort = app.readString(qs, "sort", "id")
	payload.Filters.SortSafeList = []string{"id", "name", "species", "-id", "-name", "-species"}

	if data.ValidateFilters(v, payload.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	touhous, metadata, err := app.models.Touhous.GetAll(payload.Name, payload.Species, payload.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"metadata": metadata, "touhous": touhous}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

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

func (app *application) updateTouhouHandler(w http.ResponseWriter, r *http.Request) {
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

	var payload struct {
		Name      *string  `json:"name"`
		Species   *string  `json:"species"`
		Abilities []string `json:"abilities"`
	}

	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return

	}

	if payload.Name != nil {
		touhou.Name = *payload.Name
	}

	if payload.Species != nil {
		touhou.Species = *payload.Species
	}

	if payload.Abilities != nil {
		touhou.Abilities = payload.Abilities
	}

	v := validator.New()

	if data.ValidateTouhou(v, touhou); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Touhous.Update(touhou)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"touhou": touhou}, nil)
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

	err = app.writeJSON(w, http.StatusOK, envelope{"touhou": touhou}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteTouhouHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Touhous.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)

		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "touhou successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
