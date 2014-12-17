package main

import (
	"flag"
	"log"
	"net/url"

	"./libwebcat"
)

func main() {
	flag.Parse()
	log.Println(flag.Arg(0))
	u, _ := url.Parse(flag.Arg(0))
	client, _ := libwebcat.NewWSClient(u, libwebcat.NewSTDIO())
	log.Fatal(client.Run())
}
