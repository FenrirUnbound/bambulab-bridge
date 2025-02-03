package main

import (
	"log"

	"github.com/fenrirunbound/bambulab-bridge/pkg/command"
)

func main() {
	cmd := command.NewMainCommand()

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
