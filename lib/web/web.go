package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/quiteawful/Gjallarhorn/lib/config"
)

type App struct {
	Host     string
	Port     int
	URL      string
	UseTLS   bool
	Keyfile  string
	Certfile string
	IsProxy  bool
	AssetDir string

	Mux *mux.Router
}

// New creates a new web App based on the main config
func New(cfg config.HttpdConfig) (*App, error) {
	app := &App{
		Host:     cfg.Host,
		Port:     cfg.Port,
		URL:      fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		UseTLS:   cfg.UseTLS,
		Cert:     cfg.Certfile,
		Key:      cfg.Keyfile,
		IsProxy:  cfg.InternalMode,
		AssetDir: cfg.RootDir,
		Mux:      mux.NewRouter(),
	}

	err := app.addHandlers()
	if err != nil {
		return nil, err
	}
	return app, nil
}

// Name returns the name for logging and the service manager
func (a App) Name() string { return "web" }

// Run starts the webserver
func (a App) Run() {
	// maybe TLS only with redirect
	http.ListenAndServe(a.URL, a.Mux)
}

func (a App) addHandlers() error {

	a.Mux.PathPrefix("/stati/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(a.AssetDir+"/static/"))))

	return nil
}
