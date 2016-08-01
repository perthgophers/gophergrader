package policysimulator

import (
	"encoding/json"
	"fmt"
	"github.com/blockninja/ninjarouter"
	"github.com/perthgophers/gophergrader/db"
	"github.com/perthgophers/gophergradergrading"
	"golang.org/x/net/context"
	"googlemaps.github.io/maps"
	"net/http"
)

var gmaps *maps.Client

type suburb struct {
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
}

func Serve(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
}

func getsuburb(w http.ResponseWriter, r *http.Request) {
	addr := r.FormValue("suburb")
	addr += ", WA, Australia"

	fmt.Println(addr)

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

func suburbs(w http.ResponseWriter, r *http.Request) {
	dbclient := db.Client()
	queryStr := `
		SELECT DISTINCT
		 ON (name) name,
		 description
		FROM
		 abs_mesh_data
		ORDER BY
		 name
	`
	sbs := []suburb{}

	err := dbclient.Select(&sbs, queryStr)

	jsonresponse, err := json.Marshal(sbs)
	checkErr(w, err)

	w.Write(jsonresponse)
}

func Init(gmapsclient *maps.Client) {
	gmaps = gmapsclient
	rtr := ninjarouter.New()

	rtr.GET("/suburb", getsuburb)
	rtr.GET("/suburbs", suburbs)
	rtr.GET("/*", Serve(http.FileServer(http.Dir("./policysimulator/web/public/"))))

	fmt.Println("Listening on port 9001...")
	if err := rtr.Listen(":9001"); err != nil {
		panic(err)
	}
}

func checkErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
}
