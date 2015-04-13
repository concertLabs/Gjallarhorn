package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	HttpRoot        string `json:"HttpRoot"`
	ImportDirectory string `json:"ImportDirectory"`
}

func loadConfig() Config {
	content, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var c Config
	err = json.Unmarshal(content, &c)
	if err != nil {
		panic(err)
	}
	return c
}
