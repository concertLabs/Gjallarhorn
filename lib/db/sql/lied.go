package sql

import (
	"database/sql"
	"fmt"
	"log"

	gj "github.com/quiteawful/Gjallarhorn"
)

const (
	selectLied = `
		SELECT
			id,
			titel,
			untertitel,
			jahr,
			komponist_id,
			text_id,
			arrangeur_id,
			verlag_id
		FROM
			lied 

	`
	insertLied = `
		INSERT INTO
			lied(
				titel,
				untertitel,
				jahr,
				komponist_id,
				text_id,
				arrangeur_id,
				verlag_id
			)
		VALUES ($1, $2, $3, $4, $5, $6, $7);
		`
	deleteLied = `
		DELETE FROM
			lied
		WHERE
			id = $1;
	`
)

type LiedProvider struct {
	DB *sql.DB
}

func (p *LiedProvider) Get(id int) (*gj.Lied, error) {
	var x gj.Lied

	r := p.DB.QueryRow(selectLied+" WHERE id = $1;", id)
	err := r.Scan(
		&x.ID,
		&x.Titel,
		&x.Untertitel,
		&x.Jahr,
		&x.KomponistID,
		&x.TextID,
		&x.ArrangeurID,
		&x.VerlagID,
	)

	if err != nil {
		return nil, err
	}
	return &x, nil
}

func (p *LiedProvider) GetAll() ([]*gj.Lied, error) {
	var result []*gj.Lied

	rows, err := p.DB.Query(selectLied)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var x gj.Lied

		err = rows.Scan(
			&x.ID,
			&x.Titel,
			&x.Untertitel,
			&x.Jahr,
			&x.KomponistID,
			&x.TextID,
			&x.ArrangeurID,
			&x.VerlagID,
		)

		if err != nil {
			log.Printf("error while scanning lied row: %v\n", err)
			continue
		}

		result = append(result, &x)
	}

	return result, nil
}

func (p *LiedProvider) Create(l *gj.Lied) error {
	if l.Titel == "" {
		return fmt.Errorf("titel must be set")
	}

	_, err := p.DB.Exec(
		insertLied,
		l.Titel,
		l.Untertitel,
		l.Jahr,
		l.KomponistID,
		l.TextID,
		l.ArrangeurID,
		l.VerlagID,
	)

	if err != nil {
		return err
	}
	return nil
}

func (p *LiedProvider) Delete(id int) error {
	if id <= 0 {
		return fmt.Errorf("id must be greater zero")
	}

	_, err := p.DB.Exec(deleteLied, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *LiedProvider) Search(q string) ([]*gj.Lied, error) {
	panic("not implemented")
}

func (p *LiedProvider) Edit(l *gj.Lied) error {
	panic("not implemented")
}
