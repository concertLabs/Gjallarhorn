package sql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	gj "github.com/quiteawful/Gjallarhorn"
)

const (
	selectPerson = `
		SELECT
			id,
			name,
			surname,
			street,
			zipcode,
			city,
			birth_date,
			member_since,
			email,
			password,
			access_level,
			created_at
		FROM
			person `
	insertPerson = `
		INSERT INTO
			person(
				name,
				surname,
				street,
				zipcode,
				city,
				birth_date,
				member_since,
				email,
				password,
				access_level,
				created_at
			)
			VALUES( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);

	`
)

type PersonProvider struct {
	DB *sql.DB
}

func (s *PersonProvider) Get(id int) (*gj.Person, error) {
	var x *gj.Person

	r := s.DB.QueryRow(selectPerson+" WHERE id = $1;", id)
	err := r.Scan(
		x.ID,
		x.Name,
		x.Surname,
		x.Street,
		x.Zipcode,
		x.City,
		x.BirthDate,
		x.MemberSince,
		x.Email,
		x.Password,
		x.AccessLevel,
		x.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return x, nil
}

func (s *PersonProvider) GetAll() ([]*gj.Person, error) {
	var result []*gj.Person

	rows, err := s.DB.Query(selectPerson)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var x gj.Person

		err = rows.Scan(
			&x.ID,
			&x.Name,
			&x.Surname,
			&x.Street,
			&x.Zipcode,
			&x.City,
			&x.BirthDate,
			&x.MemberSince,
			&x.Email,
			&x.Password,
			&x.AccessLevel,
			&x.CreatedAt,
		)
		if err != nil {
			log.Printf("error while scanning row: %v\n", err)
			continue
		}

		result = append(result, &x)
	}
	return result, nil
}

func (s *PersonProvider) Create(p *gj.Person) error {
	if p.Name == "" && p.Surname == "" {
		return fmt.Errorf("name and surname are required")
	}

	// currently we do not want the inserted id
	_, err := s.DB.Exec(
		insertPerson,
		p.Name,
		p.Surname,
		p.Street,
		p.Zipcode,
		p.City,
		p.BirthDate,
		p.MemberSince,
		p.Email,
		p.Password,
		p.AccessLevel,
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PersonProvider) Delete(id int) error {
	panic("not implemented")
}

func (s *PersonProvider) Search(q string) ([]*gj.Person, error) {
	panic("not implemented")
}
