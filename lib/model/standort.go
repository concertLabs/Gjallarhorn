package model

type Standort struct {
	Name    string
	Regale  []Regal
	Adresse string
	PLZ     string
	Ort     string
	Land    string
}

func (s *Standort) Save() {

}

func (s *Standort) Load() {

}
