package app

import (
	"log"

	"go.uber.org/fx"
)

func Run() {
	log.Println("Running the app")

	fx.New().Run()
}
