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
	IndexHandler   *IndexHandler
	PersonHandler  *PersonHandler
	LiedHandler    *LiedHandler
	GruppenHandler *GruppenHandler
	VerlagHandler  *VerlagHandler
}

// New creates a new web App based on the main config
func New(
	cfg config.HttpdConfig,
	ps gjallarhorn.PersonService,
	ls gjallarhorn.LiedService,
	vs gjallarhorn.VerlagService,
	gs gjallarhorn.GruppenService,
) (*App, error) {

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
	app.PersonHandler = NewPersonHandler(ps, app.Renderer)
	app.LiedHandler = NewLiedHandler(ls, ps, vs, app.Renderer)
	app.GruppenHandler = NewGruppenHandler(gs, app.Renderer)
	app.VerlagHandler = NewVerlagHandler(vs, app.Renderer)

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
	// TODO: there needs to be a huuge refactoring
	// we need a seperate api path for all this stuff

	// create each handler
	a.Mux.HandleFunc("/", a.IndexHandler.Index).Methods("GET")

	a.Mux.HandleFunc("/person", a.PersonHandler.Index).Methods("GET")
	a.Mux.HandleFunc("/person/add", a.PersonHandler.Create).Methods("GET")
	a.Mux.HandleFunc("/person/add", parseForm(a.PersonHandler.CreatePOST)).Methods("POST")
	a.Mux.HandleFunc("/person/show/{id:[0-9]+}", parseID(a.PersonHandler.Show, "/person/show/")).Methods("GET")
	a.Mux.HandleFunc("/person/delete/{id:[0-9]+}", parseID(a.PersonHandler.Delete, "/person/delete/")).Methods("GET")
	a.Mux.HandleFunc("/person/delete/{id:[0-9]+}", a.PersonHandler.DeletePOST).Methods("POST")

	a.Mux.HandleFunc("/lied", a.LiedHandler.Index).Methods("GET")
	a.Mux.HandleFunc("/lied/add", a.LiedHandler.Create).Methods("GET")
	a.Mux.HandleFunc("/lied/add", parseForm(a.LiedHandler.CreatePOST)).Methods("POST")
	a.Mux.HandleFunc("/lied/show/{id:[0-9]+}", parseID(a.LiedHandler.Show, "/lied/show/")).Methods("GET")
	a.Mux.HandleFunc("/lied/delete/{id:[0-9]+}", parseID(a.LiedHandler.Delete, "/lied/delete/")).Methods("GET")
	a.Mux.HandleFunc("/lied/delete/{id:[0-9]+}", a.LiedHandler.DeletePOST).Methods("POST")

	a.Mux.HandleFunc("/gruppe", a.GruppenHandler.Index).Methods("GET")
	a.Mux.HandleFunc("/gruppe/add", a.GruppenHandler.Create).Methods("GET")
	a.Mux.HandleFunc("/gruppe/add", a.GruppenHandler.CreatePOST).Methods("POST")
	a.Mux.HandleFunc("/gruppe/delete/{id:[0-9]+}", parseID(a.GruppenHandler.Delete, "/gruppe/delete/")).Methods("GET")
	a.Mux.HandleFunc("/gruppe/delete/{id:[0-9]+}", a.GruppenHandler.DeletePOST).Methods("POST")

	a.Mux.HandleFunc("/verlag", a.VerlagHandler.Index).Methods("GET")
	a.Mux.HandleFunc("/verlag/add", a.VerlagHandler.Create).Methods("GET")
	a.Mux.HandleFunc("/verlag/add", parseForm(a.VerlagHandler.CreatePOST)).Methods("POST")
	a.Mux.HandleFunc("/verlag/show/{id:[0-9]+}", parseID(a.VerlagHandler.Show, "/verlag/delete/")).Methods("GET")
	a.Mux.HandleFunc("/verlag/delete/{id:[0-9]+}", parseID(a.VerlagHandler.Delete, "/verlag/delete/")).Methods("GET")
	a.Mux.HandleFunc("/verlag/delete/{id:[0-9]+}", a.VerlagHandler.DeletePOST).Methods("POST")

	pp := a.Mux.PathPrefix("/static/")
	fs := http.FileServer(http.Dir(a.AssetDir + "/static/"))
	pp.Handler(http.StripPrefix("/static/", fs)) // ? how does that work again o.O

	return nil
}
