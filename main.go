package main

import (
	"fmt"
	"os"

	"github.com/viknesh-nm/sellerapp/app"
	"github.com/viknesh-nm/sellerapp/conf"
)

func main() {

	// Read Environment
	env, err := conf.Read()
	if err != nil {
		fmt.Println("Error while reading config: ", err)
		os.Exit(0)
	}

	// Initialize configuration
	conf.Init(env)

	// Initialize application
	h, err := app.Init()
	if err != nil {
		fmt.Println("Error initialising the app: ", err)
		os.Exit(1)
	}

	h.Logger.Fatal(h.Start(conf.Config.Port))

}
