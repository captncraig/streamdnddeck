package main

import (
	"fmt"
	"image/color"

	streamdeck "github.com/magicmonkey/go-streamdeck"
)

type HPPage struct {
	char *Character
}

func (h *HPPage) Render(sd *streamdeck.Device) {
	sd.WriteImageToButton(idx_Home, img("home.png"))
	sd.WriteImageToButton(0, img("heart.jpg"))
	sd.WriteTextToButton(1, fmt.Sprint(h.char.Hp.Current), color.White, color.Black)
	sd.WriteTextToButton(2, fmt.Sprint(h.char.Hp.Temp), color.Gray{Y: 150}, color.Black)
	sd.WriteTextToButton(5, "-1", red, color.Black)
	sd.WriteTextToButton(10, "-10", red, color.Black)
	sd.WriteTextToButton(6, "+1", green, color.Black)
	sd.WriteTextToButton(11, "+10", green, color.Black)
	sd.WriteTextToButton(7, "T+", green, color.Black)
	sd.WriteTextToButton(12, "T-", red, color.Black)
}
func (h *HPPage) ButtonPress(btnIndex int, sd *streamdeck.Device) bool {
	switch btnIndex {
	case 5:
		h.char.Hp.Hurt(1)
		return true
	case 6:
		h.char.Hp.Heal(1)
		return true
	case 10:
		h.char.Hp.Hurt(10)
		return true
	case 11:
		h.char.Hp.Heal(10)
		return true
	case 7:
		h.char.Hp.Temp++
		return true
	case 12:
		if h.char.Hp.Temp == 0 {
			return false
		}
		h.char.Hp.Temp--
		return true
	}
	return false
}
func (h *HPPage) ButtonRelease(btnIndex int, sd *streamdeck.Device) bool {
	switch btnIndex {
	case idx_Home:
		changePage(nil)
		return false
	}
	return false
}
func (h *HPPage) Tick() bool {
	return false
}
