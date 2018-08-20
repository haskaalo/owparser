package owparser

import (
	"errors"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Stats specific to a mode
type Stats struct {
	TimePlayed  map[string]float64                           `json:"timeplayed"`
	HeroList    []string                                     `json:"herolist"`
	CareerStats map[string]map[string]map[string]interface{} `json:"careerstats"`
}

// ErrInvalidMode The mode provided doesn't exist
var ErrInvalidMode = errors.New("This mode doesn't exist")

// TimeReplacer Replace Career profile hero time with a valid ParseDuration string format
var TimeReplacer = strings.NewReplacer(" ", "", "minutes", "m", "minute", "m", "hours", "h", "hour", "h", "seconds", "s", "second", "s")

// RegexStatName Create a better JSON-looking key name for stats
var RegexStatName = regexp.MustCompile(`[^a-zA-Z0-9]`)

// NewStats Stats specific to a mode
func (c *CareerProfile) NewStats(mode Mode) *Stats {
	stats := new(Stats)
	stats.TimePlayed = make(map[string]float64)
	stats.CareerStats = make(map[string]map[string]map[string]interface{})
	modedoc := c.document.Find("#" + string(mode)).First()

	// Time played
	modedoc.Find("div[data-category-id=\"overwatch.guid.0x0860000000000021\"] > div").Each(func(index int, selection *goquery.Selection) {
		heroTime := TimeReplacer.Replace(selection.Find(".description").First().Text())
		if heroTime == "--" {
			stats.TimePlayed[selection.Find(".title").First().Text()] = 0
			return
		}

		duration, err := time.ParseDuration(heroTime)
		if err != nil {
			log.Println(err)
			stats.TimePlayed[selection.Find(".title").First().Text()] = 0
			return
		}

		stats.TimePlayed[selection.Find(".title").First().Text()] = duration.Seconds()
	})

	// Career Stats <- Hero list
	modedoc.Find("select[data-js=\"career-select\"][data-group-id=\"stats\"] > option").Each(func(index int, selection *goquery.Selection) {
		heroName, _ := selection.Attr("option-id")
		heroID, _ := selection.Attr("value")
		stats.CareerStats[heroName] = make(map[string]map[string]interface{})

		// Hero list <- Category Name
		modedoc.Find("div[data-group-id=\"stats\"][data-category-id=\"" + heroID + "\"] > div").Each(func(index int, selection *goquery.Selection) {
			categoryName := selection.Find("h5").First().Text()
			stats.CareerStats[heroName][categoryName] = make(map[string]interface{})

			// Category Name <- Stat Name
			selection.Find("tbody > tr").Each(func(index int, selection *goquery.Selection) {
				statName := RegexStatName.ReplaceAllString(selection.Find("td:nth-child(1)").Text(), "$1")

				if strings.Contains(statName, "Time") {
					stats.CareerStats[heroName][categoryName][statName] = selection.Find("td:nth-child(2)").Text()
					return
				}
			})
		})
	})

	return stats
}
