package data

import (
	"context"
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

func (t TouhouModel) GetAll(name string, species string, filters Filters) ([]*Touhou, error) {
	query := `
		SELECT
			*
		FROM
			touhous
		ORDER BY
			id ASC`

	rows, err := t.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	touhous := []*Touhou{}

	for rows.Next() {
		var touhou Touhou
		err := rows.Scan(
			&touhou.ID,
			&touhou.CreatedAt,
			&touhou.Name,
			&touhou.Species,
			pq.Array(&touhou.Abilities),
			&touhou.Version,
		)
		if err != nil {
			return nil, err
		}
		touhous = append(touhous, &touhou)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return touhous, nil
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := t.DB.QueryRowContext(ctx, query, id).Scan(
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return t.DB.QueryRowContext(ctx, query, args...).Scan(&touhou.ID, &touhou.CreatedAt, &touhou.Version)
}

func (t TouhouModel) Update(touhou *Touhou) error {
	query := `
		UPDATE
			touhous
		SET
			name = $1, species = $2, abilities = $3, version = version + 1
		WHERE
			id = $4 AND version = $5
		RETURNING
			version`

	args := []any{touhou.Name, touhou.Species, pq.Array(touhou.Abilities), touhou.ID, touhou.Version}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := t.DB.QueryRowContext(ctx, query, args...).Scan(&touhou.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict

		default:
			return err
		}
	}
	return nil
}

func (t TouhouModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `

	DELETE FROM
		touhous
	WHERE
		id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	result, err := t.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
