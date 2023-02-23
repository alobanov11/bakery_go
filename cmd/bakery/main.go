package main

import (
	"bakery/internal/pkg/app"
	"flag"
	"log"
)

var (
	t = flag.Int("t", 8, "number of working time in seconds")
	c = flag.Int("c", 3, "number of concurrent tasks")
	f = flag.String("f", "etc/bakes.json", "path to bakes.json")
)

func main() {
	flag.Parse()

	app, err := app.New(*t, *c, *f)

	if err != nil {
		log.Fatal(err)
	}

	app.Start()
}
