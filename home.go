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
	idxAttacks    = 8
	idxStats      = 9
	idxChecks     = 10
	idxSaves      = 11
	idxAC         = 12
	idxInv        = 13
)

var red = color.RGBA{R: 255, G: 0, B: 0, A: 255}
var green = color.RGBA{R: 0, G: 255, B: 0, A: 255}

func img(path string) string {
	return filepath.Join("images", path)
}

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
	return img
}

func (h *homePage) Render(sd *streamdeck.Device) {
	log.Printf("Rendering home for %s", h.char.Name)
	sd.WriteImageToButton(idxIcon, img(h.char.Icon))

	sd.WriteRawImageToButton(idxHp, renderHpButton(h.char))

	sd.WriteRawImageToButton(idxSpellSlots, renderSSButton(h.char))
	sd.WriteImageToButton(idxDice, img("d20.jpg"))
	sd.WriteImageToButton(idxDaughter, img("wolf.jpg"))

	shield, err := getImageFromFile("shield.png")
	if err != nil {
		log.Fatal(err)
	}
	drawTextToImage(fmt.Sprint(h.char.AC), color.Black, shield, 40, 50)
	sd.WriteRawImageToButton(idxAC, shield)

	sorc, err := getImageFromFile("sorc.png")
	if err != nil {
		log.Fatal(err)
	}
	drawTextToImage(fmt.Sprint(h.char.SorcPoints), color.Black, sorc, 20, 50)
	sd.WriteRawImageToButton(idxSorc, sorc)

	sd.WriteImageToButton(idxRest, img("sleep.png"))
	sd.WriteImageToButton(idxAttacks, img("sword.jpg"))
	sd.WriteImageToButton(idxInv, img("bag.jpg"))
	sd.WriteImageToButton(idxResources, img("resources.jpg"))
}
func (h *homePage) ButtonPress(btnIndex int, sd *streamdeck.Device) bool {
	switch btnIndex {
	case idxAC:
		// toggle ac between 13 and 16 for mage armor
		if h.char.AC == 13 {
			h.char.AC = 16
		} else {
			h.char.AC = 13
		}
		return true
	}
	return false

}
func (h *homePage) ButtonRelease(btnIndex int, sd *streamdeck.Device) bool {
	switch btnIndex {
	case idxDice:
		changePage(&rollPage{char: h.char})
		return false
	}
	return false
}
func (h *homePage) Tick() bool {
	return false
}
