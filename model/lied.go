package noten

type Lied struct {
	Titel     string
	Komponist string
	Genre     string
	Verlag    string
	Stimmen   []Stimme
}

func (l *Lied) Save() {

}

func (l *Lied) Load() {

}
