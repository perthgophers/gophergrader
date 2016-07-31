package community

import (
	"fmt"
	"strconv"

	"github.com/perthgophers/govhack/db"
)

type CommunityLocation struct {
	Rank string `db:"community_rank"`
}

// Location returns an integer score of community
func Location(longitude, latitude float64) (int, error) {
	fmt.Println(longitude, latitude)
	longStr := strconv.FormatFloat(longitude, 'f', 6, 64)
	latStr := strconv.FormatFloat(latitude, 'f', 6, 64)
	dbclient := db.Client()
	queryStr := fmt.Sprintf(`
       SELECT 

        case 
            when count(distinct feature_cl||geographic) < 2 then 1 
            when count(distinct feature_cl||geographic) < 5 then 4 
            when count(distinct feature_cl||geographic) < 10 then 5 
            when count(distinct feature_cl||geographic) < 15 then 7 
            when count(distinct feature_cl||geographic) < 20 then 9 
        else 10 

        end as community_rank 
        FROM geonoma 
        WHERE feature_cl in('Monument', 'Marine Park', 'Park', 'Pool', 'Recreation', 'Golf Course', 'Community') AND 
        ST_DWithin(geometry, ST_GeomFromText('POINT(%s %s)',4326), 0.01) 
        `, longStr, latStr)
	fmt.Println(queryStr)
	communityLocations := []CommunityLocation{}
	err := dbclient.Select(&communityLocations, queryStr)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	result, err := strconv.Atoi(communityLocations[0].Rank)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println("Community Location Score:", result)
	return result, nil
}
