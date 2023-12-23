package main

import (
	"fmt"
	"github.com/LuizSSampaio/gopportunities/config"
	"github.com/LuizSSampaio/gopportunities/internal/router"
)

func main() {
	// Initialize the config
	configErr := config.Init()
	if configErr != nil {
		fmt.Println(configErr)
		return
	}

	// Initialize the router
	router.Init()
}
