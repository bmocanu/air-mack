package main

import (
	"math/rand"
	"time"
)

func main() {
	// config.init() 	=> sets up the configuration based on env vars
	// retrieve.init() 	=> sets up the cron tab and schedules the periodical comic scanning
	// store.init() 	=> sets up the issue storage
	// serve.init()		=> sets up the HTTP interface + controllers
	rand.Seed(time.Now().UnixNano())

	// initial retrieving upon startup
	//retrieve.FetchNewIssues()

	// start listening for incoming HTTP requests
	//serve.StartServing()
}
