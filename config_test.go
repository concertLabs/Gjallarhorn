package main

import "testing"

func TestLoadConfig(t *testing.T) {
	c, err := loadConfig("config.json")
	if err != nil {
		t.Fatalf("Error while loading Config file: %s\n", err.Error())
	}

	if c.Httpd.InternalMode != true {
		t.Errorf("Error.")
	}
}
