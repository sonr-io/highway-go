/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"context"
	"log"

	"github.com/sonr-io/highway-go/config"
	highway "github.com/sonr-io/highway-go/server"
)

func main() {

	// This loads the information from the configuration file into a struct
	cnfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// This is starting the Highway based Node utilizing the configuration
	err = highway.Start(context.Background(), cnfg)
	if err != nil {
		log.Fatal(err)
	}

	// This is bad practice however this is being used to block the code
	// from completing execution
	//select {}
}
