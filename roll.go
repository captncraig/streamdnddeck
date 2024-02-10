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
	idx_Roll_Info   = 5
	idx_Roll_D20    = 1
	idx_Roll_Adv    = 2
	idx_Roll_Dis    = 3
	idx_Roll_D100   = 4
	idx_Roll_D12    = 6
	idx_Roll_D10    = 7
	idx_Roll_D8     = 8
	idx_Roll_D6     = 9
	idx_Roll_D4     = 11
	idx_Roll_D2     = 12
)

type rollPage struct {
	accum     int
	rolling   bool
	rollType  int
	rollCount int
	advType   string

	resultA int
	resultB int
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
		if h.advType == "A" {
			sd.WriteTextToButton(idx_Roll_Result, fmt.Sprint(h.resultB), green, color.Black)
			sd.WriteTextToButton(idx_Roll_Info, fmt.Sprint(h.resultA), color.Gray{Y: 150}, color.Black)
		} else if h.advType == "D" {
			sd.WriteTextToButton(idx_Roll_Result, fmt.Sprint(h.resultA), red, color.Black)
			sd.WriteTextToButton(idx_Roll_Info, fmt.Sprint(h.resultB), color.Gray{Y: 150}, color.Black)
		} else {
			sd.WriteTextToButton(idx_Roll_Result, fmt.Sprint(h.resultA), color.White, color.Black)
			if h.rollCount > 0 {
				img := newSolid(color.Black)
				txt := fmt.Sprintf("%dd%d", h.rollCount, h.rollType)
				drawTextToImage(txt, color.White, img, 18, 40)
				sd.WriteRawImageToButton(idx_Roll_Info, img)
			}
		}
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
		sd.WriteImageToButton(idx_Roll_D4, img("caltrop.png"))
		sd.WriteImageToButton(idx_Roll_D2, img("coin.jpg"))
	}
}

func (h *rollPage) ButtonPress(btnIndex int, sd *streamdeck.Device) bool {
	if h.rolling {
		return false
	}
	startRoll := func(i int, adv string) bool {
		h.rolling = true
		h.advType = adv
		if h.rollType == i {
			h.rollCount++
		} else {
			h.rollCount = 1
			h.accum = 0
		}
		h.rollType = i
		return true
	}
	switch btnIndex {
	case idx_Roll_Accum:
		h.rollType = 0
		h.rollCount = 0
		h.accum = 0
		h.advType = ""
		return true
	case idx_Roll_D20:
		return startRoll(20, "")
	case idx_Roll_D100:
		return startRoll(100, "")
	case idx_Roll_D12:
		return startRoll(12, "")
	case idx_Roll_D10:
		return startRoll(10, "")
	case idx_Roll_D8:
		return startRoll(8, "")
	case idx_Roll_D6:
		return startRoll(6, "")
	case idx_Roll_D4:
		return startRoll(4, "")
	case idx_Roll_D2:
		return startRoll(2, "")
	case idx_Roll_Adv:
		return startRoll(20, "A")
	case idx_Roll_Dis:
		return startRoll(20, "D")
	}
	return false
}
func (h *rollPage) ButtonRelease(btnIndex int, sd *streamdeck.Device) bool {
	endRoll := func(i int, adv string) bool {
		if i != h.rollType || adv != h.advType {
			return false
		}
		h.rolling = false
		h.resultA = rand.Intn(h.rollType) + 1
		if h.advType != "" {
			h.resultB = rand.Intn(h.rollType) + 1
			// swap so a <= b
			if h.resultB < h.resultA {
				x := h.resultA
				h.resultA = h.resultB
				h.resultB = x
			}
		}
		if h.rollType != 20 && h.rollType != 100 {
			h.accum += h.resultA
		}
		return true
	}
	if h.rolling {
		switch btnIndex {
		case idx_Roll_D20:
			return endRoll(20, "")
		case idx_Roll_D100:
			return endRoll(100, "")
		case idx_Roll_D12:
			return endRoll(12, "")
		case idx_Roll_D10:
			return endRoll(10, "")
		case idx_Roll_D8:
			return endRoll(8, "")
		case idx_Roll_D6:
			return endRoll(6, "")
		case idx_Roll_D4:
			return endRoll(4, "")
		case idx_Roll_D2:
			return endRoll(2, "")
		case idx_Roll_Adv:
			return endRoll(20, "A")
		case idx_Roll_Dis:
			return endRoll(20, "D")
		}
	} else {
		switch btnIndex {
		case idx_Home:
			changePage(nil)
			return false
		}
	}
	return false
}
func (h *rollPage) Tick() bool {
	return h.rolling
}
