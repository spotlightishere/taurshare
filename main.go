package main

import (
	_ "github.com/sijms/go-ora/v2"
)

func main() {
	// Ensure we have loaded our config.
	loadConfig()
	// Bring up the database.
	loadDatabase()
}
