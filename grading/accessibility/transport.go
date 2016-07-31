package accessibility

import (
	"fmt"
	"strconv"

	"github.com/perthgophers/govhack/db"
)

type TransportResult struct {
	Rank string `db:"pta_rank"`
}

// Transport returns a slice of TransportResults
func Transport(longitude, latitude float64) (int, error) {
	fmt.Println(longitude, latitude)
	longStr := strconv.FormatFloat(longitude, 'f', 6, 64)
	latStr := strconv.FormatFloat(latitude, 'f', 6, 64)
	dbclient := db.Client()
	queryStr := fmt.Sprintf(`
        SELECT 

        case 
            when count(name) <1 then 1 
            when count(name) <2 then 2 
            when count(name) <3 then 3 
            when count(name) <4 then 4 
            when count(name) <5 then 5 
            when count(name) <6 then 6 
            when count(name) <7 then 7 
            when count(name) <8 then 8 
            when count(name) <9 then 9 
        else 10 

        end as pta_rank 

        FROM "public transport authority stops _ bus stops" 
        where ST_DWithin(wkb_geometry, ST_GeomFromText('POINT(%s %s)',4326), 0.003) 
        `, longStr, latStr)
	fmt.Println(queryStr)
	transports := []TransportResult{}
	err := dbclient.Select(&transports, queryStr)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	result, err := strconv.Atoi(transports[0].Rank)
	if err != nil {
		return 0, err
	}
	return result, nil
}
