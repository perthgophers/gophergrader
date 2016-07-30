package grading

import (
	"github.com/perthgophers/govhack/grading/safety"
	"googlemaps.github.io/maps"
)

//Cafe 1km
func Grade(addr []maps.GeocodingResult) int {
	longitude := addr[0].Geometry.Location.Lng
	latitude := addr[0].Geometry.Location.Lat
	hospitalscore, _ := safety.Hospitals(longitude, latitude)
	return hospitalscore
}
