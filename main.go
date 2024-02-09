package main

import (
	"fmt"
	"log"
	"os"

	streamdeck "github.com/magicmonkey/go-streamdeck"
	_ "github.com/magicmonkey/go-streamdeck/devices"
	"gopkg.in/yaml.v3"
)

type Character struct {
	Name       string `yaml:"name"`
	Icon       string `yaml:"icon"`
	Hp         HPInfo `yaml:"hp"`
	SpellSlots SSInfo `yaml:"spell_slots"`
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
	ButtonPress(btnIndex int, sd *streamdeck.Device)
	ButtonRelease(btnIndex int, sd *streamdeck.Device)
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

	sd.ButtonPress(func(btnIndex int, sd *streamdeck.Device, err error) {
		fmt.Println(btnIndex)
		if err != nil {
			panic(err)
		}
		currentPage.ButtonPress(btnIndex, sd)
	})

	sd.ButtonRelease(func(btnIndex int, sd *streamdeck.Device, err error) {
		fmt.Println("!!!", btnIndex)
		currentPage.ButtonRelease(btnIndex, sd)
	})

	select {}
}
