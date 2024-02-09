package main

import (
	"fmt"
	"log"
	"os"
	"time"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
	"gopkg.in/yaml.v3"
)

type Character struct {
	Name          string `yaml:"name"`
	Icon          string `yaml:"icon"`
	Hp            HPInfo `yaml:"hp"`
	SpellSlots    SSInfo `yaml:"spell_slots"`
	AC            int    `yaml:"ac"`
	SorcPoints    int    `yaml:"sorc_points"`
	SorcPointsMax int    `yaml:"sorc_points_max"`
}
type HPInfo struct {
	Current int `yaml:"current"`
	Max     int `yaml:"max"`
	Temp    int `yaml:"temp"`
	MaxMod  int `yaml:"max_mod"`
}
type SSInfo struct {
	Current []int `yaml:"current"`
	Max     []int `yaml:"max"`
}

type Page interface {
	Render(*streamdeck.Device)
	ButtonPress(btnIndex int, sd *streamdeck.Device) bool
	ButtonRelease(btnIndex int, sd *streamdeck.Device) bool
	Tick() bool
}

var newPage = make(chan Page, 100)

func changePage(p Page) {
	newPage <- p
}

func main() {
	sd, err := streamdeck.Open()
	if err != nil {
		panic(err)
	}
	fmt.Println(sd.GetName())
	sd.ClearButtons()
	sd.SetBrightness(50)

	charBytes, err := os.ReadFile("chars/hyacinth.yaml")
	if err != nil {
		log.Fatal(err)
	}
	hyacinth := &Character{}
	if err = yaml.Unmarshal(charBytes, hyacinth); err != nil {
		log.Fatal(err)
	}
	fmt.Println(hyacinth)
	var currentPage Page = &homePage{char: hyacinth}

	currentPage.Render(sd)
	down := make(chan int, 100)
	up := make(chan int, 100)

	sd.ButtonPress(func(btnIndex int, sd *streamdeck.Device, err error) {
		if err != nil {
			panic(err)
		}
		down <- btnIndex
	})

	sd.ButtonRelease(func(btnIndex int, sd *streamdeck.Device, err error) {
		if err != nil {
			panic(err)
		}
		up <- btnIndex
	})

	tick := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case p := <-newPage:
			currentPage = p
			sd.ClearButtons()
			p.Render(sd)
		case <-tick.C:
			if currentPage.Tick() {
				currentPage.Render(sd)
			}
		case i := <-down:
			if currentPage.ButtonPress(i, sd) {
				currentPage.Render(sd)
			}
		case i := <-up:
			if currentPage.ButtonRelease(i, sd) {
				currentPage.Render(sd)
			}
		}

	}
}
