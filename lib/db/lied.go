package db

type Lied struct {
	Titel       string
	Untertitel  string
	Jahr        int
	Komponist   Person
	KomponistID int
	Texter      Person
	TexterID    int
	Arrangeur   Person
	ArrangeurID int
	Verlag      Verlag
	VerlagID    int
}
