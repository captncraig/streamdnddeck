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

const filename = "chars/hyacinth.yaml"

func loadChar() *Character {
	charBytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	hyacinth := &Character{}
	if err = yaml.Unmarshal(charBytes, hyacinth); err != nil {
		log.Fatal(err)
	}
	fmt.Println(hyacinth)
	return hyacinth
}

func saveChar(c *Character) {
	dat, err := yaml.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filename, dat, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	sd, err := streamdeck.Open()
	if err != nil {
		panic(err)
	}
	fmt.Println(sd.GetName())
	sd.ClearButtons()
	sd.SetBrightness(50)

	var currentPage Page = &homePage{char: loadChar()}
	home := currentPage
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
	btnsDown := map[int]bool{}
	for {
		select {
		case p := <-newPage:
			if p == nil {
				p = home
			}
			btnsDown = map[int]bool{}
			currentPage = p
			sd.ClearButtons()
			p.Render(sd)
		case <-tick.C:
			if currentPage.Tick() {
				currentPage.Render(sd)
			}
		case i := <-down:
			btnsDown[i] = true
			if currentPage.ButtonPress(i, sd) {
				currentPage.Render(sd)
			}
		case i := <-up:
			// only send release event if there was a down since the last state change
			if !btnsDown[i] {
				continue
			}
			btnsDown[i] = false
			if currentPage.ButtonRelease(i, sd) {
				currentPage.Render(sd)
			}
		}

	}
}
