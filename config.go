package main

import (
	"encoding/xml"
	"io/ioutil"
)

// Config represents credentials necessary to connect to ATP.
type Config struct {
	ConnectionString string
	WalletPassword   string
	DatabaseUser     string
	DatabasePassword string
}

// config is a global way to reference the program's current config.
var config Config

// loadConfig parses credentials as specified within the user's config.
func loadConfig() {
	contents, err := ioutil.ReadFile("config.xml")
	if err != nil {
		panic(err)
	}

	err = xml.Unmarshal(contents, &config)
	if err != nil {
		panic(err)
	}
}
