package libwebcat

import "flag"

var Origin string

func init() {
	flag.StringVar(&Origin, "o", "", "Setup the Origin Header")
}
