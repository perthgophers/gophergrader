package grading

import (
	"fmt"
	"math/rand"

	"github.com/perthgophers/govhack/grading/accessibility"
	"github.com/perthgophers/govhack/grading/community"
	"github.com/perthgophers/govhack/grading/safety"
	"googlemaps.github.io/maps"
)

type GradingResult struct {
	Accessibility int `json:"accessibility"`
	Apocalypse    int `json:"apocalypse"`
	Community     int `json:"community"`
	Culture       int `json:"culture"`
	Safety        int `json:"safety"`
	Services      int `json:"services"`
}

//Cafe 1km
func Grade(addr []maps.GeocodingResult) GradingResult {
	longitude := addr[0].Geometry.Location.Lng
	latitude := addr[0].Geometry.Location.Lat
	safety.Hospitals(longitude, latitude)
	accessibiltyTransportScore, err := accessibility.Transport(longitude, latitude)
	if err != nil {
		fmt.Println(err)
	}

	communityLocationScore, _ := community.Location(longitude, latitude)
	fmt.Println("Community Location Score:", communityLocationScore)
	fmt.Println("Accessibility Transport Score:", accessibiltyTransportScore)
	results := GradingResult{
		Accessibility: accessibiltyTransportScore,
		Apocalypse:    rand.Intn(7) + 3,
		Community:     rand.Intn(7) + 3,
		Culture:       rand.Intn(7) + 3,
		Safety:        rand.Intn(7) + 3,
		Services:      rand.Intn(7) + 3,
	}

	return results
}
