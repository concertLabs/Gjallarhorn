package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quiteawful/Gjallarhorn/api/v100"
)

type WebApp struct {
	RootDir string
	Mux     *mux.Router
}

func NewWebApp(root string) WebApp {
	return WebApp{RootDir: root, Mux: mux.NewRouter()}
}

func (app *WebApp) Run() {
	app.Mux.HandleFunc("/", app.IndexHandler)
	app.Mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(app.RootDir+"/static/"))))
	a100 := api100.GetSubrouter("/api/v100")
	app.Mux.PathPrefix("/api/v100").Handler(a100)

	http.ListenAndServe(":8080", app.Mux)
}

// IndexHandler is the main page
func (app *WebApp) IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := app.loadTemplate("base", "index")
	if err != nil {
		log.Printf("[Template] %v\n", err)
		return
	}
	t.Execute(w, nil)
}
