package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"path/filepath"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

type homePage struct {
	char *Character
}

const (
	idxIcon       = 0
	idxHp         = 1
	idxSpellSlots = 2
	idxSorc       = 3
	idxDice       = 4
	idxResources  = 5
	idxRest       = 6
	idxDaughter   = 7
)

var red = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var green = color.RGBA{R: 0, G: 255, B: 0, A: 255}

func renderHpButton(c *Character) image.Image {
	hp, err := getImageFromFile("heart-corner.png")
	if err != nil {
		log.Fatal(err)
	}
	drawTextToImage(fmt.Sprint(c.Hp.Current), color.White, hp, 45, 60)
	drawTextToImage("Hp", red, hp, 18, 15)
	return hp
}

func renderSSButton(c *Character) image.Image {
	img := newSolid(color.Black)
	ss := c.SpellSlots
	log.Println(ss)
	off := -5
	xoff := 8
	split := 72 / 3
	i := 0
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			txt := fmt.Sprint(ss.Current[i])
			var col color.Color = color.White
			if ss.Max[i] == 0 {
				continue
			}
			if ss.Current[i] == 0 {
				col = color.Gray{Y: 50}
			} else if ss.Current[i] < ss.Max[i] {
				col = color.Gray{Y: 150}
			}
			drawTextToImageX(txt, col, img, 18, (split*(y+1))+off, split*x+xoff)
			i++
		}
	}
	//drawTextToImage("1 0 1", red, i, 18, split+off)
	//drawTextToImage("1 0 1", red, i, 18, split+off+split)
	//drawTextToImage("1 0 1", red, i, 18, split+off+split+split)
	return img
}

func (h *homePage) Render(sd *streamdeck.Device) {
	log.Printf("Rendering home for %s", h.char.Name)
	sd.WriteImageToButton(idxIcon, filepath.Join("images", h.char.Icon))

	sd.WriteRawImageToButton(idxHp, renderHpButton(h.char))

	sd.WriteRawImageToButton(idxSpellSlots, renderSSButton(h.char))
	sd.WriteImageToButton(idxDice, filepath.Join("images", "d20.jpg"))
	sd.WriteImageToButton(idxDaughter, filepath.Join("images", "wolf.jpg"))

}
func (h *homePage) ButtonPress(btnIndex int, sd *streamdeck.Device) {

}
func (h *homePage) ButtonRelease(btnIndex int, sd *streamdeck.Device) {

}
