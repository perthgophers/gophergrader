package safety

import (
	"fmt"
	"github.com/perthgophers/govhack/db"
	"strconv"
)

type HospitalResult struct {
	Descriptino string `db:"geographic"`
	Category    string `db:"feature_cl"`
	Geometry    string `db:"geometry"`
}

func Hospitals(longitude, latitude float64) (int, error) {
	fmt.Println(longitude, latitude)

	longStr := strconv.FormatFloat(longitude, 'f', 6, 64)
	latStr := strconv.FormatFloat(latitude, 'f', 6, 64)

	dbclient := db.Client()

	hospitals := []HospitalResult{}

	queryStr := `
		SELECT *
		FROM geonoma
		WHERE feature_cl = 'Health' AND
		ST_DWithin(
	    Geography(ST_MakePoint(?, ?)),
	    ?
	  );
	`

	fmt.Println(queryStr)

	err := dbclient.Select(&hospitals, queryStr, longStr, latStr, 2)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	for k, i := range hospitals {
		fmt.Println(k, i)
	}

	return 0, nil
}
