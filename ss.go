package main

import (
	"fmt"
	"image/color"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

type SSPage struct {
	char   *Character
	adding bool
}

var ords = []string{"1st:", "2nd:", "3rd:", "4th:", "5th:", "6th:", "7th:", "8th", "9th"}

func (h *SSPage) Render(sd *streamdeck.Device) {
	sd.ClearButtons()
	sd.WriteImageToButton(0, img("wand.jpg"))
	sd.WriteImageToButton(idx_Home, img("home.png"))
	for i := 1; i <= 9; i++ {
		img := newSolid(color.Black)

		drawTextToImageX(ords[i-1], color.White, img, 18, 15, 5)
		drawTextToImageX(fmt.Sprint(h.char.SpellSlots.Current[i-1]), color.White, img, 30, 50, 10)
		drawTextToImageX(fmt.Sprintf("/%d", h.char.SpellSlots.Max[i-1]), color.White, img, 25, 55, 35)
		sd.WriteRawImageToButton(i, img)
	}
	sd.WriteTextToButton(13, "+", color.White, color.Black)
}
func (h *SSPage) ButtonPress(btnIndex int, sd *streamdeck.Device) bool {
	if btnIndex >= 1 && btnIndex <= 9 {
		if h.adding {
			if h.char.SpellSlots.Current[btnIndex-1] < h.char.SpellSlots.Max[btnIndex-1] {
				h.char.SpellSlots.Current[btnIndex-1]++
				return true
			}
		} else {
			if h.char.SpellSlots.Current[btnIndex-1] > 0 {
				h.char.SpellSlots.Current[btnIndex-1]--
				return true
			}
		}
	} else if btnIndex == 13 {
		h.adding = true
		return false
	}
	return false
}
func (h *SSPage) ButtonRelease(btnIndex int, sd *streamdeck.Device) bool {
	switch btnIndex {
	case idx_Home:
		changePage(nil)
		return false
	case 13:
		h.adding = false
		return false
	}
	return false
}
func (h *SSPage) Tick() bool {
	return false
}
