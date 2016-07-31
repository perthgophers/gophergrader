package accessibility

import (
	"fmt"
	"github.com/perthgophers/govhack/db"
	"strconv"
)

type CongestionResult struct {
	MeanValue float64 `db:"trafficrank"`
}

func Congestion(longitude, latitude float64) (int, error) {
	score := []CongestionResult{};

	longStr := strconv.FormatFloat(longitude, 'f', 6, 64)
	latStr := strconv.FormatFloat(latitude, 'f', 6, 64)

	/*
	queryStr := `
		SELECT has_congestion_mean
		FROM traffic_congestion_rank
		WHERE
		poly_line not like '%)' AND
		ST_DWithin(ST_GeomFromText(poly_line), ST_GeomFromText($1,4326), 1000)
	`
	*/
	queryStr := `
		SELECT 
		case 
			when avg(has_congestion_mean) <  0.06321347 then 1 
			when avg(has_congestion_mean) <  0.103310502 then 2 
			when avg(has_congestion_mean) <  0.1293379 then 3 
			when avg(has_congestion_mean) <  0.146261416 then 4 
			when avg(has_congestion_mean) <  0.173401827 then 5 
			when avg(has_congestion_mean) <  0.206478311 then 6 
			when avg(has_congestion_mean) <  0.242037671 then 7 
			when avg(has_congestion_mean) <  0.292751142 then 8 
			when avg(has_congestion_mean) <  0.363527397 then 9 
			else 10 
		end	AS trafficrank
		FROM traffic_congestion_rank 
		WHERE 
		poly_line like '%)' AND NOT poly_line LIKE 'MULTI%' AND ST_DWithin(ST_GeomFromText(poly_line, 4326), ST_GeomFromText($1, 4326), 0.5);
	`

	dbclient := db.Client()

	err := dbclient.Select(
		&score,
		queryStr,
		fmt.Sprintf("POINT(%s %s)", longStr, latStr),
	)

	//fmt.Printf("phor %+v\n", score)

	/*
	hashmap := make(map[float64]float64){
		0.063213: 1,
		0.103311: 2,
		0.129338: 3,
		0.146261: 4,
		0.173402: 5,
		0.206478: 6,
		0.242038: 7,
		0.292751: 8,
		0.363527: 9,
	}
	*/

	if err != nil {
		fmt.Println("err", err)
		return 0, err
	}

	for k, i := range score {
		fmt.Println(k, i)
	}

	if len(score) >=1 {
		return int(score[0].MeanValue), nil
	}
	return 0, nil
}
