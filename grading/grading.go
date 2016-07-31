package grading

import (
	"github.com/perthgophers/govhack/grading/safety"
	"googlemaps.github.io/maps"
	"math/rand"
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

	results := GradingResult{
		Accessibility: rand.Intn(10),
		Apocalypse:    rand.Intn(10),
		Community:     rand.Intn(10),
		Culture:       rand.Intn(10),
		Safety:        rand.Intn(10),
		Services:      rand.Intn(10),
	}

	return results
}
