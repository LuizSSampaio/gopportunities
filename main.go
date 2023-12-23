package main

import (
	"github.com/LuizSSampaio/gopportunities/config"
	"github.com/LuizSSampaio/gopportunities/internal/router"
)

func main() {
	logger := config.GetLogger()

	// Initialize the config
	configErr := config.Init()
	if configErr != nil {
		logger.Fatalf("Error initializing config: %s", configErr.Error())
		return
	}

	// Initialize the router
	router.Init()
}
