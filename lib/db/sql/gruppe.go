package sql

import (
	"database/sql"
	"fmt"
	"log"

	gj "github.com/quiteawful/Gjallarhorn"
)

const (
	selectGruppe = `
		SELECT
			id,
			name
		FROM
			gruppe 
	`
	insertGruppe = `
		INSERT INTO
			gruppe (
				name
			)
		VALUES ($1);
	`
	updateGruppe = `
		UPDATE
			gruppe
		SET
			name = $1
		WHERE
			id = $2
	`
	deleteGruppe = `
		DELETE FROM 
			gruppe
		WHERE

			id = $1;
	`
)

type GruppenProvider struct {
	DB *sql.DB
}

func (g *GruppenProvider) Get(id int) (*gj.Gruppe, error) {
	var x gj.Gruppe

	r := g.DB.QueryRow(selectGruppe+" WHERE id = $1;", id)
	err := r.Scan(
		&x.ID,
		&x.Name,
	)

	if err != nil {
		return nil, err
	}
	return &x, nil
}

func (g *GruppenProvider) GetAll() ([]*gj.Gruppe, error) {
	var result []*gj.Gruppe

	rows, err := g.DB.Query(selectGruppe)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var x gj.Gruppe

		err = rows.Scan(
			&x.ID,
			&x.Name,
		)

		if err != nil {
			log.Printf("error while scanning grupen row: %v\n", err)
			continue
		}

		result = append(result, &x)
	}

	return result, nil
}

func (g *GruppenProvider) Create(gr *gj.Gruppe) error {
	if gr.Name == "" {
		return fmt.Errorf("name must be set")
	}

	_, err := g.DB.Exec(
		insertGruppe,
		gr.Name,
	)

	if err != nil {
		return err
	}
	return nil
}

func (g *GruppenProvider) Edit(gr *gj.Gruppe) error {
	if gr.Name == "" {
		return fmt.Errorf("name must not be empty")
	}

	_, err := g.DB.Exec(
		updateGruppe,
		gr.Name,
	)

	if err != nil {
		return err
	}
	return nil
}

func (g *GruppenProvider) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("id must be greater zero")
	}

	_, err := g.DB.Exec(
		deleteGruppe,
		id,
	)

	if err != nil {
		return err
	}
	return nil
}
