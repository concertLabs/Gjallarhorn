package main

import (
	"flag"
	"log"

	"github.com/quiteawful/Gjallarhorn/lib/config"
)

type args struct {
	command string
	config  string
}

const (
	cmdNothing = "noop"
	cmdDB      = "db"
)

func parseArgs() args {
	var r args

	flag.StringVar(&r.command, "cmd", cmdNothing, "noop: do nothing, db: init database")
	flag.StringVar(&r.config, "config", "config.json", "path to the json formatted config file, default: config.json")

	flag.Parse()
}

func main() {
	args := parseArgs()
	if args.command == cmdNothing {
		return
	}

	cfg, err := config.Open(args.config)
	if err != nil {
		log.Fatalf("could not open config file: %v\n", err)
	}

	switch args.command {
	case cmdDB:
		log.Printf("init database")
		break
	}

}
