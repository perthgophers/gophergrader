package grading

import (
	"fmt"

	"github.com/perthgophers/govhack/grading/accessibility"
	"github.com/perthgophers/govhack/grading/apocalypse"
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
	accessibiltyBusScore, _ := accessibility.Bus(longitude, latitude)
	communityLocationScore, _ := community.Location(longitude, latitude)
	accessibilityCongestionScore, _ := accessibility.Congestion(longitude, latitude)
	safetyFinalScore, _ := safety.Crime(longitude, latitude)

	congestionScore, _ := accessibility.Congestion(longitude, latitude)
	rainFallScore, _ := apocalypse.RainFall(longitude, latitude)

	fmt.Println("congi", congestionScore)
	fmt.Println("raifc", rainFallScore)
	accessibiltyFinalScore := int((float64(accessibiltyBusScore) + float64(accessibilityCongestionScore)) / 2.0)
	communityFinalScore := communityLocationScore
	servicesFinalScore, _ := community.Service(longitude, latitude)
	fmt.Println("Accessibility Final Score:", accessibiltyFinalScore)
	fmt.Println("Community Final Score:", communityFinalScore)
	fmt.Println("Safety Final Score:", safetyFinalScore)
	fmt.Println("Services Final Score:", servicesFinalScore)

	results := GradingResult{
		Accessibility: accessibiltyFinalScore,
		Apocalypse:    5,
		Community:     communityFinalScore,
		Culture:       5,
		Safety:        safetyFinalScore,
		Services:      servicesFinalScore,
	}
	return results
}
