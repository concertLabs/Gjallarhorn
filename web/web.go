package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quiteawful/Gjallarhorn/api/v100"
	"github.com/quiteawful/Gjallarhorn/lib/config"
)

// WebApp is our web gui
type WebApp struct {
	RootDir string
	Host    string
	Port    int
	IsProxy bool
	Mux     *mux.Router
}

// NewWebApp returns a new instance of our webapp
func NewWebApp(cfg config.HttpdConfig) WebApp {
	return WebApp{
		RootDir: cfg.AssetDir,
		Host:    cfg.Host,
		Port:    cfg.Port,
		IsProxy: cfg.InternalMode,
		Mux:     mux.NewRouter(),
	}
}

// Name returns the name of this module
func (app WebApp) Name() string {
	return "web"
}

// Run is the main run func
func (app WebApp) Run() {
	app.Mux.HandleFunc("/", app.IndexHandler)
	app.Mux.HandleFunc("/lied", app.LiedIndexHandler)

	app.Mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(app.RootDir+"/static/"))))
	a100 := api100.GetSubrouter("/api/v100")
	app.Mux.PathPrefix("/api/v100").Handler(a100)

	url := fmt.Sprintf("%s:%d", app.Host, app.Port)
	log.Printf("[WebApp] Start httpd on %s\n", url)
	http.ListenAndServe(url, app.Mux)
}

// IndexHandler is the main page
func (app WebApp) IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := app.loadTemplate("base", "index")
	if err != nil {
		log.Printf("[Template] %v\n", err)
		return
	}
	t.Execute(w, nil)
}
