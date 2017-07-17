package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config represents the json structure of the configuration file for Gjallarhorn
// each part of the programm has its own sub-struct
type Config struct {
	Httpd    HttpdConfig  `json:"Httpd"`
	Importer ImportConfig `json:"Importer"`
	API      APIConfig    `json:"Api"`
}

// HttpdConfig can be used as our main http frontend server with special configs
// or we can sit behind a nginx/apache server
type HttpdConfig struct {
	// InternalMode decides wether to use own httpd or use
	// other server to serve html content, eg. nginx
	InternalMode bool   `json:"InternalMode"`
	RootDir      string `json:"RootDir"`
}

// ImportConfig controls the automatic importer to scan the filesystem for new pdf files
// and imports them in the database
type ImportConfig struct {
	// Folder to watch for new files
	ScanDir     string `json:"ScanDir"`
	UseImporter bool   `json:"UseImporter"`
}

// APIConfig I do not know what this does :) phil!!f
type APIConfig struct {
	Port int `json:"Port"`
}

func loadConfig(file string) (*Config, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var c *Config
	err = json.Unmarshal(content, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
