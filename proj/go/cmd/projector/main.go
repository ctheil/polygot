package main

import (
	"fmt"
	"log"

	"calebtheil.com/polygot/pkg/config"
)

func main() {
	opts, err := config.GetOpts()
	if err != nil {
		log.Fatalf("Unable to get options: %v", err)
	}
	config, err := config.NewConfig(opts)
	if err != nil {
		log.Fatalf("Unable to get config: %v", err)
	}
	fmt.Printf("confing: %+v", config)
}
