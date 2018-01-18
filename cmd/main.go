package main

import "github.com/jgimeno/avatarme/src"

func main() {
	i := identicon.Generate("pedo")

	identicon.Render(i)
}
