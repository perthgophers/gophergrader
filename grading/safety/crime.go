package safety

import (
	"fmt"
	"strconv"

	"github.com/perthgophers/gophergrader/db"
)

type CrimeRank struct {
	Rank string `db:"offences_per_capita_rank"`
}

// Crime returns a slice of CrimeRankss
func Crime(longitude, latitude float64) (int, error) {
	fmt.Println(longitude, latitude)
	longStr := strconv.FormatFloat(longitude, 'f', 6, 64)
	latStr := strconv.FormatFloat(latitude, 'f', 6, 64)
	dbclient := db.Client()
	queryStr := fmt.Sprintf(`
        select 

case 

when CAST(offences.sum as DECIMAL(10, 3)) / CAST(crime_population.sum as DECIMAL(10, 3)) <1 then 10 
when CAST(offences.sum as DECIMAL(10, 3)) / CAST(crime_population.sum as DECIMAL(10, 3)) <2 then 9 
when CAST(offences.sum as DECIMAL(10, 3)) / CAST(crime_population.sum as DECIMAL(10, 3)) <3 then 8 
when CAST(offences.sum as DECIMAL(10, 3)) / CAST(crime_population.sum as DECIMAL(10, 3)) <4 then 7 
when CAST(offences.sum as DECIMAL(10, 3)) / CAST(crime_population.sum as DECIMAL(10, 3)) <5 then 6 
when CAST(offences.sum as DECIMAL(10, 3)) / CAST(crime_population.sum as DECIMAL(10, 3)) <6 then 5 
when CAST(offences.sum as DECIMAL(10, 3)) / CAST(crime_population.sum as DECIMAL(10, 3)) <7 then 4 
when CAST(offences.sum as DECIMAL(10, 3)) / CAST(crime_population.sum as DECIMAL(10, 3)) <8 then 3 
when CAST(offences.sum as DECIMAL(10, 3)) / CAST(crime_population.sum as DECIMAL(10, 3)) <9 then 2 
when CAST(offences.sum as DECIMAL(10, 3)) / CAST(crime_population.sum as DECIMAL(10, 3)) <999 then 1 
else 5 


end as offences_per_capita_rank 




from crime_population 

join 

                (select * 

                from crime_offences 

                join (SELECT rtrim(split_part(abs_mesh_data.name, ' ', 1)) as name FROM abs_mesh_data WHERE ST_Within(ST_GeomFromText('POINT(%s %s)', 4326), wkb_geometry) LIMIT 1) as suburb 

                on crime_offences.key ilike suburb.name) as offences 

on crime_population.key = offences.name 
        `, longStr, latStr)
	fmt.Println(queryStr)
	crimes := []CrimeRank{}
	err := dbclient.Select(&crimes, queryStr)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	if len(crimes) == 0 {
		return 0, nil
	}
	result, err := strconv.Atoi(crimes[0].Rank)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return result, nil
}
