package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Httpd    HttpdConfig  `json:"Httpd"`
	Importer ImportConfig `json:"Importer"`
	Api      ApiConfig    `json:"Api"`
}

type HttpdConfig struct {
	// InternalMode decides wether to use own httpd or use
	// other server to serve html content, eg. nginx
	InternalMode bool   `json:"InternalMode"`
	RootDir      string `json:"RootDir"`
}

type ImportConfig struct {
	// Folder to watch for new files
	ScanDir     string `json:"ScanDir"`
	UseImporter bool   `json:"UseImporter"`
}

type ApiConfig struct {
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
