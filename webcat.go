package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/thypon/webcat/libwebcat"
)

func main() {
	flag.Parse()
	log.Println(flag.Arg(0))
	u, err := url.Parse(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	client, err := libwebcat.NewWSClient(u, libwebcat.NewSTDIO())
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(client.Run())
}
