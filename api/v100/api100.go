package api100

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quiteawful/Gjallarhorn/db/v100"
	"github.com/quiteawful/Gjallarhorn/noten"
)

func GetSubrouter(prefix string) *mux.Router {
	r := mux.NewRouter().PathPrefix(prefix).Subrouter()
	//Lied
	r.HandleFunc("/lied", postLiedHandler).Methods("POST")
	r.HandleFunc("/lied/{name}", getLiedHandler).Methods("GET")
	r.HandleFunc("/lied/{name}/addStimme", addStimmetoLiedHandler).Methods("POST")
	//Stimme
	r.HandleFunc("/lied/{name}", postStimmeHandler).Methods("POST")
	r.HandleFunc("/lied/{liedname}/{stimmenname}", getStimmeHandler).Methods("GET")
	//Repertoire
	r.HandleFunc("/repertoire", postRepertoireHandler).Methods("POST")
	r.HandleFunc("/repertoire/{name}", getRepertoireHandler).Methods("GET")
	r.HandleFunc("/repertoire/{name}/addLied", addLiedtoRepertoirHandler).Methods("POST")
	//Schublade
	r.HandleFunc("/schublade", postSchubladeHandler).Methods("POST")
	r.HandleFunc("/schublade/{name}", getSchubladeHandler).Methods("GET")
	r.HandleFunc("/schublade/{name}/addLied", addLiedtoSchubladeHandler).Methods("POST")
	//Regal
	r.HandleFunc("/regal", postRegalHandler).Methods("POST")
	r.HandleFunc("/regal/{name}", getRegalHandler).Methods("GET")
	r.HandleFunc("/regal/{name}/addSchublade", addSchubladetoRegalHandler).Methods("POST")
	//Standort
	r.HandleFunc("/standort", postStandortHandler).Methods("POST")
	r.HandleFunc("/standort/{name}", getStandortHandler).Methods("GET")
	r.HandleFunc("/standort/{name}/addRegal", addRegaltoStandortHandler).Methods("POST")

	return r
}

func getLiedHandler(w http.ResponseWriter, r *http.Request) {

}

func postLiedHandler(w http.ResponseWriter, r *http.Request) {
	var l noten.Lied
	l.Titel = r.Header.Get("Titel")
	l.Komponist = r.Header.Get("Komponist")
	l.Genre = r.Header.Get("Genre")
	l.Verlag = r.Header.Get("Verlag")

	rev, err := db100.DB.Save(l, l.Titel+l.Komponist, "")
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, rev)
}

func addStimmetoLiedHandler(w http.ResponseWriter, r *http.Request) {

}

func getStimmeHandler(w http.ResponseWriter, r *http.Request) {

}

func postStimmeHandler(w http.ResponseWriter, r *http.Request) {

}

func getRepertoireHandler(w http.ResponseWriter, r *http.Request) {

}

func postRepertoireHandler(w http.ResponseWriter, r *http.Request) {

}

func addLiedtoRepertoirHandler(w http.ResponseWriter, r *http.Request) {

}

func getSchubladeHandler(w http.ResponseWriter, r *http.Request) {

}

func postSchubladeHandler(w http.ResponseWriter, r *http.Request) {

}

func addLiedtoSchubladeHandler(w http.ResponseWriter, r *http.Request) {

}

func getRegalHandler(w http.ResponseWriter, r *http.Request) {

}

func postRegalHandler(w http.ResponseWriter, r *http.Request) {

}

func addSchubladetoRegalHandler(w http.ResponseWriter, r *http.Request) {

}

func getStandortHandler(w http.ResponseWriter, r *http.Request) {

}

func postStandortHandler(w http.ResponseWriter, r *http.Request) {

}

func addRegaltoStandortHandler(w http.ResponseWriter, r *http.Request) {

}
