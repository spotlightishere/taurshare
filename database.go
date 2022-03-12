package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"regexp"

	_ "github.com/sijms/go-ora/v2"
)

var (
	ErrInvalidTNS = errors.New("unable to parse given TNS connection string")
)

var db *sql.DB

// loadDatabase instantiates a connection with our remote database.
func loadDatabase() {
	// Parse the given TNS connection string to obtain a DSN.
	dsn := parseTNS()
	fmt.Println(dsn)

	// Perform our connection.
	var err error
	db, err = sql.Open("oracle", dsn)

	// Confirm our connection was successful.
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	defer db.Close()
}

// tnsRegex parses a port, hostname, and service name from a TNS connection string.
var tnsRegex = regexp.MustCompile(`.*\(port=(\d*)\)\(host=(.*)\)\).*\(service_name=(.*oraclecloud.com)\)\).*`)

// parseTNS transforms a TNS connection string into a usable DSN.
// It additionally specifies user credentials via configuration.
func parseTNS() string {
	// Prepare our DSN.
	var dsn url.URL
	dsn.Scheme = "oracle"
	dsn.User = url.UserPassword(config.DatabaseUser, config.DatabasePassword)

	// Obtain a port, hostname, and service name from our TNS connection string.
	matches := tnsRegex.FindStringSubmatch(config.ConnectionString)
	if matches == nil {
		panic(ErrInvalidTNS)
	}

	// Obtain values from our matches.
	// Note that the port is matched first, and the host immediately afterwards.
	port := matches[1]
	host := matches[2]
	serviceName := matches[3]

	// Update our DSN accordingly.
	dsn.Host = fmt.Sprintf("%s:%s", host, port)
	dsn.Path = serviceName

	// We want SSL, and verify based on our wallet.
	// We additionally specify our wallet path.
	var query = make(url.Values)
	query.Set("SSL", "true")
	query.Set("SSL Verify", "enable")
	query.Set("Wallet", "./wallet")
	dsn.RawQuery = query.Encode()

	return dsn.String()
}
