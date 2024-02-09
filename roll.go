package main

import (
	"fmt"
	"image/color"
	"math/rand"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

const (
	idx_Home        = 14
	idx_Roll_Accum  = 10
	idx_Roll_Result = 0
	idx_Roll_D20    = 1
	idx_Roll_Adv    = 2
	idx_Roll_Dis    = 3
	idx_Roll_D100   = 4
	idx_Roll_D12    = 6
	idx_Roll_D10    = 7
	idx_Roll_D8     = 8
	idx_Roll_D6     = 9
)

type rollPage struct {
	char     *Character
	accum    int
	rolling  bool
	rollType int
	advType  string
}

func (h *rollPage) Render(sd *streamdeck.Device) {
	if h.rolling {
		for i := 0; i <= 14; i++ {
			if i == idx_Roll_Accum {
				continue
			}
			sd.WriteColorToButton(i, color.RGBA{
				R: uint8(rand.Intn(255)),
				G: uint8(rand.Intn(255)),
				B: uint8(rand.Intn(255)),
				A: 255,
			})
		}
	} else {
		sd.ClearButtons()
		sd.WriteImageToButton(idx_Home, img("home.png"))
		sd.WriteTextToButton(idx_Roll_Accum, fmt.Sprint(h.accum), color.White, color.Black)
		sd.WriteImageToButton(idx_Roll_D20, img("d20.jpg"))
		sd.WriteImageToButton(idx_Roll_Adv, img("thumb-up.png"))
		sd.WriteImageToButton(idx_Roll_Dis, img("thumb-down.png"))
		sd.WriteTextToButton(idx_Roll_D100, "100", color.White, color.Black)
		sd.WriteImageToButton(idx_Roll_D12, img("d12.png"))
		sd.WriteImageToButton(idx_Roll_D10, img("d10.png"))
		sd.WriteImageToButton(idx_Roll_D8, img("d8.png"))
		sd.WriteImageToButton(idx_Roll_D6, img("d6.png"))
	}
}

func (h *rollPage) ButtonPress(btnIndex int, sd *streamdeck.Device) bool {
	switch btnIndex {
	case idx_Roll_D20:
		if h.rolling {
			return false
		}
		h.rolling = true
		h.rollType = 20
		return true
	}
	return false
}
func (h *rollPage) ButtonRelease(btnIndex int, sd *streamdeck.Device) bool {
	switch btnIndex {
	case idx_Home:
		changePage(&homePage{char: h.char})
		return false
	case idx_Roll_D20:
		h.rolling = false
		return true
	}
	return false
}
func (h *rollPage) Tick() bool {
	return h.rolling
}
