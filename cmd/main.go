package main

import (
	"github.com/jgimeno/avatarme/src"
	"flag"
)

var baseString = flag.String("string", "baseString", "Base string to generate the Identicon.")

func main() {
	flag.Parse()
	i := identicon.Generate(*baseString)
	identicon.Render(i)
}
