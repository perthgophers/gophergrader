package apocalypse

import (
	"fmt"
	"strconv"

	"github.com/perthgophers/govhack/db"
)

type RainFallResult struct {
	Rank      float64 `db:"rainfall_rank"`
	MeanValue float64 `db:"value_mean"`
}

func RainFall(longitude, latitude float64) (int, error) {
	fmt.Println(longitude, latitude)

	longStr := strconv.FormatFloat(longitude, 'f', 6, 64)
	latStr := strconv.FormatFloat(latitude, 'f', 6, 64)

	dbclient := db.Client()

	rainFallResult := []RainFallResult{}

	/*
		queryStr := `
			SELECT value_mean
			FROM bom_avg_monthly_rainfall
		    WHERE
		    ST_Distance(ST_GeomFromText($1,4326), geometry) > 0
		`
	*/
	queryStr := `
		SELECT 
			value_mean, 
			case 
				when value_mean<54.9 then        1 
				when value_mean<134.2 then        2 
				when value_mean<280.6 then        3 
				when value_mean<390.4 then        4 
				when value_mean<616.1 then        5 
				when value_mean<1061.4 then        6 
				when value_mean<1622.6 then        7 
				when value_mean<2366.8 then        8 
				when value_mean<3446.5 then        9 
			else 10 
			end as rainfall_rank, 
			ST_Distance(geometry, ST_GeomFromText($1,4326)) AS distant 

	    FROM  bom_avg_monthly_rainfall 
	    WHERE 
	    	value_mean>60 AND
	    	ST_Distance(geometry, ST_GeomFromText($1,4326)) < 500 
	    ORDER BY distant 
	    LIMIT 1 
	`
	err := dbclient.Select(
		&rainFallResult,
		queryStr,
		fmt.Sprintf("POINT(%s %s)", longStr, latStr),
	)

	fmt.Println("rfrqi", len(rainFallResult))

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	for k, i := range rainFallResult {
		fmt.Println("rf", k, i)
	}

	if len(rainFallResult) >= 1 {
		return int(rainFallResult[0].MeanValue), nil
	}
	return 0, nil
}
