package data

import (
	"database/sql"
	"errors"
	"time"

	"github.com/arvindeva/touhouapi-cms/internal/validator"
	"github.com/lib/pq"
)

// Touhou represents a Touhou character.
type Touhou struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"_"`
	Name      string    `json:"name"`
	Species   string    `json:"species"`
	Abilities []string  `json:"abilities"`
	Version   int32     `json:"version"`
}

type TouhouModel struct {
	DB *sql.DB
}

func ValidateTouhou(v *validator.Validator, touhou *Touhou) {
	v.Check(touhou.Name != "", "name", "must be provided")
	v.Check(len(touhou.Name) < 500, "name", "must not be more than 500 bytes")

	v.Check(touhou.Species != "", "species", "must be provided")

	v.Check(validator.Unique(touhou.Abilities), "abilities", "must not contain duplicate values")
}

func (t TouhouModel) Get(id int64) (*Touhou, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT
			id, created_at, name, species, abilities, version
		FROM
			touhous
		WHERE
			id = $1`

	var touhou Touhou
	err := t.DB.QueryRow(query, id).Scan(
		&touhou.ID,
		&touhou.CreatedAt,
		&touhou.Name,
		&touhou.Species,
		pq.Array(&touhou.Abilities),
		&touhou.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &touhou, nil
}

func (t TouhouModel) Insert(touhou *Touhou) error {
	query := `
		INSERT INTO
			touhous (name, species, abilities)
		VALUES
			($1, $2, $3)
		RETURNING
			id, created_at, version`

	args := []any{touhou.Name, touhou.Species, pq.Array(touhou.Abilities)}
	return t.DB.QueryRow(query, args...).Scan(&touhou.ID, &touhou.CreatedAt, &touhou.Version)
}

func (t TouhouModel) Update(touhou *Touhou) error {
	return nil
}

func (t TouhouModel) Delete(id int64) error {
	return nil
}
