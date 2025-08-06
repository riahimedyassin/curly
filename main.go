package main

import (
	"log"

	"github.com/riahimedyassin/curly/cmd"
)

func main() {
	rootCmd := cmd.NewRootCMD()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
