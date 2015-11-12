package web

import (
	"io"
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

	a100 := api100.GetSubrouter("/api/v100")
	app.Mux.PathPrefix("/api/v100").Handler(a100)

	http.ListenAndServe(":8080", app.Mux)
}

func (app *WebApp) IndexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hi")
}
