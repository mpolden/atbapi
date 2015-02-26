package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/martinp/atbapi/atb"
	"log"
	"os"
)

func main() {
	var opts struct {
		Config string `short:"c" long:"config" description:"Path to config file" value-name:"FILE" default:"config.json"`
	}
	_, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		os.Exit(1)
	}

	client, err := atb.NewFromConfig(opts.Config)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
}