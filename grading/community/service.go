// SELECT

// case
// when count(distinct feature_cl||geographic) < 1 then 1
// when count(distinct feature_cl||geographic) < 3 then 2
// when count(distinct feature_cl||geographic) < 4 then 3
// when count(distinct feature_cl||geographic) < 5 then 4
// when count(distinct feature_cl||geographic) < 7 then 5
// when count(distinct feature_cl||geographic) < 8 then 6
// when count(distinct feature_cl||geographic) < 10 then 7
// when count(distinct feature_cl||geographic) < 15 then 8
// when count(distinct feature_cl||geographic) < 25 then 9
// else 10
// end as service_rank

//         FROM geonoma
//         WHERE feature_cl in('School','Place of Worship','Child Care Centres','Health','Western Australian Police Facility') AND
//         ST_DWithin(geometry, ST_GeomFromText('POINT(115.875170 -32.246220)',4326), 0.01)

package community

import (
	"fmt"
	"strconv"

	"github.com/perthgophers/gophergrader/db"
)

type ServiceLocation struct {
	Rank string `db:"service_rank"`
}

// Service returns an integer score of service
func Service(longitude, latitude float64) (int, error) {
	fmt.Println(longitude, latitude)
	longStr := strconv.FormatFloat(longitude, 'f', 6, 64)
	latStr := strconv.FormatFloat(latitude, 'f', 6, 64)
	dbclient := db.Client()
	queryStr := fmt.Sprintf(`
       SELECT 
        case 
        when count(distinct feature_cl||geographic) < 1 then 1 
        when count(distinct feature_cl||geographic) < 3 then 2 
        when count(distinct feature_cl||geographic) < 4 then 3 
        when count(distinct feature_cl||geographic) < 5 then 4 
        when count(distinct feature_cl||geographic) < 7 then 5 
        when count(distinct feature_cl||geographic) < 8 then 6 
        when count(distinct feature_cl||geographic) < 10 then 7 
        when count(distinct feature_cl||geographic) < 15 then 8 
        when count(distinct feature_cl||geographic) < 25 then 9 
        else 10 
        end as service_rank 

        FROM geonoma 
        WHERE feature_cl in('School','Place of Worship','Child Care Centres','Health','Western Australian Police Facility') AND 
        ST_DWithin(geometry, ST_GeomFromText('POINT(%s %s)',4326), 0.01)  
        `, longStr, latStr)
	fmt.Println(queryStr)
	serviceLocations := []ServiceLocation{}
	err := dbclient.Select(&serviceLocations, queryStr)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	if len(serviceLocations) == 0 {
		return 0, nil
	}
	result, err := strconv.Atoi(serviceLocations[0].Rank)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return result, nil
}
