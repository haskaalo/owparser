package owparser

import (
	"strconv"
	"strings"
)

// General Data available even if account is private
type General struct {
	Private  bool   `json:"private"`
	SR       int    `json:"rank,omitempty"`
	Prestige int    `json:"prestige"`
	Level    string `json:"level"`
	Portrait string `json:"portrait"`
}

var prestigeReplacer = strings.NewReplacer("background-image:url(https://d1u1mce87gyfbn.cloudfront.net/game/playerlevelrewards/", "", "_Border.png)", "")

// NewGeneral Get Data available even if account is private
func (c *CareerProfile) NewGeneral() *General {
	general := new(General)
	permission := c.document.Find(".masthead-permission-level-text").First().Text()
	if permission == "Public Profile" {
		general.Private = false
	} else {
		general.Private = true
	}

	general.SR, _ = strconv.Atoi(c.document.Find(".masthead-player .competitive-rank").First().Text())

	level := c.document.Find(".player-level").First()
	levelStyle, _ := level.Attr("style")

	general.Prestige = Prestige[prestigeReplacer.Replace(levelStyle)]
	general.Level = level.Text()

	general.Portrait, _ = c.document.Find(".player-portrait").First().Attr("src")

	return general
}
