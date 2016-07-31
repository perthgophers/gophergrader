package grading

import (
	"fmt"
	"math/rand"
	"github.com/perthgophers/govhack/grading/safety"
	"github.com/perthgophers/govhack/grading/accessibility"
	"github.com/perthgophers/govhack/grading/apocalypse"
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

	congestionScore, _ := accessibility.Congestion(longitude, latitude)
	rainFallScore, _ := apocalypse.RainFall(longitude, latitude)

	fmt.Println("congi", congestionScore)
	fmt.Println("raifc", rainFallScore)

	results := GradingResult{
		Accessibility: rand.Intn(7) + 3,
		Apocalypse:    rand.Intn(7) + 3,
		Community:     rand.Intn(7) + 3,
		Culture:       rand.Intn(7) + 3,
		Safety:        rand.Intn(7) + 3,
		Services:      rand.Intn(7) + 3,
	}

	return results
}
