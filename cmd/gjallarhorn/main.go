package main

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/quiteawful/Gjallarhorn/lib/config"
	"github.com/quiteawful/Gjallarhorn/lib/importer"
	"github.com/quiteawful/Gjallarhorn/lib/servicemanager"
	"github.com/quiteawful/Gjallarhorn/web"
)

// Args is a struct to parse the commandline arguments
type Args struct {
	config string
}

// DefaultOptions returns a struct with the minimal/default values
// might be extendet later on
func defaultArgs() Args {
	return Args{config: "config.json"}
}

func parseArgs() Args {
	var result = defaultArgs()

	flag.StringVar(&result.config, "config", result.config, "the json formatted config file")
	flag.Parse()

	return result
}

func main() {
	args := parseArgs()
	c, err := config.Open(args.config)
	if err != nil {
		log.Fatalf("could not open configfile: %v\n", err)
	}

	Importer := importer.New(c.Importer)
	WebApp := web.NewWebApp(c.Httpd.RootDir)

	// TODO: make a service manager
	// go Importer.Run()
	// WebApp.Run()

	manager := servicemanager.NewManager(2)

	manager.Add(Importer)
	manager.Add(WebApp)
	manager.Start()
	defer manager.Stop()

	// catch ctrl-c and prevent programm to stop immediatly
	ctrl := make(chan os.Signal, 1)
	ch := make(chan int)
	signal.Notify(ctrl, os.Interrupt)
	go func() {
		for {
			select {
			case <-ctrl:
				os.Exit(0)
			case <-ch:
				continue
			}
		}
	}()
	<-ch
}
