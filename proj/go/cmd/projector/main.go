package main

import (
	"encoding/json"
	"fmt"
	"log"

	"calebtheil.com/polygot/pkg/config"
)

func logError(err error) {
	log.Fatalf("Error: %e", err)
}

func main() {
	opts, err := config.GetOpts()
	if err != nil {
		log.Fatalf("Unable to get options: %v", err)
	}
	c, err := config.NewConfig(opts)
	if err != nil {
		log.Fatalf("Unable to get config: %v", err)
	}

	p := config.NewProjector(c)

	if c.Operation == config.Print {
		if len(c.Args) == 0 {
			data := p.GetValueAll()
			jsonsString, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("This line should never be reached %e", err)
			}

			fmt.Printf("%+v", string(jsonsString))
		} else {
			val, ok := p.GetValue(c.Args[0])
			if !ok {
				fmt.Printf("Could not find value for %v", c.Args[0])
			}
			fmt.Printf("%+v", val)
		}
	}
	if c.Operation == config.Add {
		p.SetValue(c.Args[0], c.Args[1])
		if err := p.Save(); err != nil {
			logError(err)
		}
	}

	if c.Operation == config.Remove {
		p.RemoveValue(c.Args[0])
		if err := p.Save(); err != nil {
			logError(err)
		}
	}
}
