package apocalypse

import (
	"fmt"
	"github.com/perthgophers/govhack/db"
	"strconv"
)

type RainFallResult struct {
	MeanValue float64 `db:"value_mean"`
}

func RainFall(longitude, latitude float64) (int, error) {
	fmt.Println(longitude, latitude)

	longStr := strconv.FormatFloat(longitude, 'f', 6, 64)
	latStr := strconv.FormatFloat(latitude, 'f', 6, 64)

	dbclient := db.Client()

	rainFallResult := []RainFallResult{}

	queryStr := `
		SELECT value_mean
		FROM bom_avg_monthly_rainfall
	    WHERE
	    ST_Distance(ST_GeomFromText($1,4326), geometry) > 0
	`

	fmt.Println("qstr", queryStr)

	err := dbclient.Select(
	    &rainFallResult,
	    queryStr,
		fmt.Sprintf("POINT(%s %s)", longStr, latStr),
	)

	fmt.Println("rfrq", rainFallResult)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	for k, i := range rainFallResult {
		fmt.Println(k, i)
	}

	return 0, nil
}
