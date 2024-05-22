package main

import (
	"pinares/cmd/router"
	"pinares/config"
)

func main() {
	e := router.New()

	err := e.Start(":" + configuration.Config_.Port)
	if err != nil {
		return
	}
}
