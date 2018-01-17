package identicon

import (
	"image"
	"os"
	"image/png"
)

func Render(i Identicon) {
	img := image.NewRGBA(image.Rect(0,0, 50, 50))

	f, _ := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	png.Encode(f, img)
}
