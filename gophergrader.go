package main

import (
	"encoding/json"
	"fmt"
	"github.com/blockninja/ninjarouter"
	"github.com/perthgophers/govhack/grading"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"net/http"
)

var gmaps *maps.Client

func Serve(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

func grade(w http.ResponseWriter, r *http.Request) {
	addr := r.FormValue("address")
	georeq := &maps.GeocodingRequest{
		Address: addr,
	}
	resp, err := gmaps.Geocode(context.Background(), georeq)
	checkErr(w, err)

	grresult := grading.Grade(resp)

	jsonresponse, err := json.Marshal(grresult)

	checkErr(w, err)

	w.Write(jsonresponse)
}

func main() {
	var err error
	gmaps, err = maps.NewClient(maps.WithAPIKey(GMapsAPIKey))
	if err != nil {
		panic(err)
	}

	rtr := ninjarouter.New()
	rtr.GET("/grade", grade)
	rtr.GET("/*", Serve(http.FileServer(http.Dir("./web/public/"))))
	fmt.Println("Listening on port 9000...")
	if err := rtr.Listen(":9000"); err != nil {
		panic(err)
	}
}

func checkErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}
