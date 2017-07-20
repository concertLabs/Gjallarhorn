package model

type Lied struct {
	Titel     string
	Komponist string
	Genre     string
	Verlag    string
	Stimmen   []Stimme
}

func NewLied() *Lied {
	return &Lied{}
}

func (l *Lied) Save() {

}

func (l *Lied) Load() {

}

