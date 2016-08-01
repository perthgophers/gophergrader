package safety

import (
	"fmt"
	"github.com/perthgophers/gophergrader/db"
	"strconv"
)

type HospitalResult struct {
	Descriptino string `db:"geographic"`
	Category    string `db:"feature_cl"`
}

func Hospitals(longitude, latitude float64) (int, error) {
	fmt.Println(longitude, latitude)

	longStr := strconv.FormatFloat(longitude, 'f', 6, 64)
	latStr := strconv.FormatFloat(latitude, 'f', 6, 64)

	dbclient := db.Client()

	hospitals := []HospitalResult{}

	queryStr := `
		SELECT geographic, feature_cl
		FROM geonoma
		WHERE feature_cl = 'Health' AND
		geographic ILIKE '%hospital%' AND
		ST_DWithin(geometry, ST_GeomFromText($1,4326), 0.1)
	`

	fmt.Println(queryStr)

	longlatstr := fmt.Sprintf("POINT(%s %s)", longStr, latStr)

	err := dbclient.Select(&hospitals, queryStr, longlatstr)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	for k, i := range hospitals {
		fmt.Println(k, i)
	}

	return 0, nil
}
