package main

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
	"path/filepath"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gomono"
)

func newSolid(c color.Color) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, 72, 72))
	// colour := color.RGBA{red, green, blue, 0}
	draw.Draw(img, img.Bounds(), image.NewUniform(c), image.Point{0, 0}, draw.Src)
	return img
}

func getImageFromFile(filename string) (image.Image, error) {
	f, err := os.Open(filepath.Join("images", filename))
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}
	return img, nil
}
func drawTextToImageX(text string, textColour color.Color, dstImage image.Image, size int, y int, x int) {

	myfont, err := truetype.Parse(gomono.TTF)
	if err != nil {
		panic(err)
	}

	srcImg := image.NewUniform(textColour)
	dimg, ok := dstImage.(draw.Image)
	if !ok {
		log.Fatal("bad image type?")
	}
	c := freetype.NewContext()
	c.SetFont(myfont)
	c.SetDst(dimg)
	c.SetSrc(srcImg)

	fs := float64(size)

	c.SetFontSize(fs)
	c.SetClip(dstImage.Bounds())

	pt := freetype.Pt(x, y)
	c.DrawString(text, pt)
}
func drawTextToImage(text string, textColour color.Color, dstImage image.Image, size int, y int) {
	btnSize := 72
	wid := getTextWidth(text, float64(size))
	x := btnSize/2 - wid/2
	drawTextToImageX(text, textColour, dstImage, size, y, x)
}

func getTextWidth(text string, size float64) int {

	myfont, err := truetype.Parse(gomono.TTF)
	if err != nil {
		panic(err)
	}

	// Calculate width of string
	width := 0
	face := truetype.NewFace(myfont, &truetype.Options{Size: size})
	for _, x := range text {
		awidth, _ := face.GlyphAdvance(rune(x))
		iwidthf := int(float64(awidth) / 64)
		width += iwidthf
	}

	return width
}
