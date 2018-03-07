package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	gjallarhorn "github.com/quiteawful/Gjallarhorn"
	"github.com/quiteawful/Gjallarhorn/lib/config"
)

// NOTE: can be deleted??? dupe of lib/web/web.go

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

	Renderer *Renderer

	// Handler with Routes
	IndexHandler  *IndexHandler
	PersonHandler *PersonHandler
}

// New creates a new web App based on the main config
func New(cfg config.HttpdConfig, personService gjallarhorn.PersonService) (*App, error) {
	app := &App{
		Host:     cfg.Host,
		Port:     cfg.Port,
		URL:      fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		UseTLS:   cfg.UseTLS,
		Certfile: cfg.Certfile,
		Keyfile:  cfg.Keyfile,
		IsProxy:  cfg.InternalMode,
		AssetDir: cfg.AssetDir,
		Mux:      mux.NewRouter(),

		Renderer: NewRenderer(cfg.AssetDir),
	}

	app.IndexHandler = NewIndexHandler(app.Renderer)
	app.PersonHandler = NewPersonHandler(personService, app.Renderer)
	// app.LiederHandler= NewLiederHandler(app.Renderer)

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
	// TODO: add custom logger to App struct
	log.Printf("Start httpd on %s\n", a.URL)
	if a.UseTLS {
		// maybe TLS only with redirect
		log.Fatal(http.ListenAndServeTLS(a.URL, a.Certfile, a.Keyfile, a.Mux))
	} else {
		log.Fatal(http.ListenAndServe(a.URL, a.Mux))
	}
}

func (a App) addHandlers() error {
	// NOTE: there will never be returned an error, might change func signature
	// TODO: there needs to be a huuge refactoring
	// we need a seperate api path for all this stuff

	// create each handler
	a.Mux.HandleFunc("/", a.IndexHandler.Index).Methods("GET")
	a.Mux.HandleFunc("/person", a.PersonHandler.Index).Methods("GET")                 // show all
	a.Mux.HandleFunc("/person/add", a.PersonHandler.CreateGET).Methods("GET")         // add new
	a.Mux.HandleFunc("/person/add", a.PersonHandler.CreatePOST).Methods("POST")       // add new
	a.Mux.HandleFunc("/person/show/{id:[0-9]+}", a.PersonHandler.Show).Methods("GET") // show a single person
	a.Mux.HandleFunc("/person/delete/{id:[0-9]+}", a.PersonHandler.DeleteGET).Methods("GET")
	a.Mux.HandleFunc("/person/delete/{id:[0-9]+}", a.PersonHandler.DeletePOST).Methods("POST")

	// just static stuff, for reading
	pp := a.Mux.PathPrefix("/static/")
	fs := http.FileServer(http.Dir(a.AssetDir + "/static/"))
	pp.Handler(http.StripPrefix("/static/", fs)) // ? how does that work again o.O

	return nil
}
