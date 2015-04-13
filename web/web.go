package web

import (
	"io"
	"net/http"
)

type WebApp struct {
	RootDir string
	Mux     *http.ServeMux
}

func NewWebApp(root string) WebApp {
	return WebApp{RootDir: root, Mux: http.NewServeMux()}
}

func (app *WebApp) Run() {
	app.Mux.HandleFunc("/", app.IndexHandler)

	http.ListenAndServe(":8080", app.Mux)
}

func (app *WebApp) IndexHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hi")
}
