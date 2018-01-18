package main

import (
	"flag"
	"github.com/jgimeno/go-identicon/src"
)

var baseString = flag.String("string", "baseString", "Base string to generate the Identicon.")

func main() {
	flag.Parse()
	i := identicon.Generate(*baseString)
	identicon.Render(i)
}
