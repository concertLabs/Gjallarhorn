package model

// Lied is our entity representation of a sheet of paper with notes
type Lied struct {
	// Titel is the common name of a song
	Titel     string
	Komponist string
	Genre     string
	Verlag    string
	Stimmen   []Stimme
	Notizen   []Notiz
}

func NewLied() *Lied {
	return &Lied{}
}

func (l *Lied) Save() {

}

func (l *Lied) Load() {

}
