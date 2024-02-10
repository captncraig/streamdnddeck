package main

type Character struct {
	Name          string `yaml:"name"`
	Icon          string `yaml:"icon"`
	Hp            HPInfo `yaml:"hp"`
	SpellSlots    SSInfo `yaml:"spell_slots"`
	AC            int    `yaml:"ac"`
	SorcPoints    int    `yaml:"sorc_points"`
	SorcPointsMax int    `yaml:"sorc_points_max"`
}

func (c *Character) LongRest() {
	c.Hp.Current = c.Hp.Max
	copy(c.SpellSlots.Current, c.SpellSlots.Max)
	c.SorcPoints = c.SorcPointsMax
	c.AC = 13
}

func (c *Character) ShortRest() {

}

type HPInfo struct {
	Current int `yaml:"current"`
	Max     int `yaml:"max"`
	Temp    int `yaml:"temp"`
	MaxMod  int `yaml:"max_mod"`
}

func (h *HPInfo) Heal(amt int) {
	for amt > 0 {
		if h.Current == (h.Max + h.MaxMod) {
			return
		}
		amt--
		h.Current++
	}
}
func (h *HPInfo) Hurt(amt int) {
	for amt > 0 {
		if h.Temp > 0 {
			h.Temp--
			amt--
			continue
		}
		if h.Current == 0 {
			return
		}
		h.Current--
		amt--
	}

}

type SSInfo struct {
	Current []int `yaml:"current"`
	Max     []int `yaml:"max"`
}
