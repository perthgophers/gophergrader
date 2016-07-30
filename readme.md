# The Gopher Grader

The Gopher Grade provides users with an address based grade using a spider graph at a point in time.

![](http://www.economicmodeling.com/wp-content/uploads/radar-graph-3.JPG)

## Project Spinup

```
$ go get github.com/perthgophers/govhack
$ glide install
$ cd web && npm run build && cd ..
$ docker run --name perthgophers-postgis -p 5432:5432 -e POSTGRES_PASSWORD=iamagopher -d mdillon/postgis
```

## Metrics

-  Culture
-  Apocalypse Readiness
-  Accessibility
-  Commerce
-  Safety

## Datasets

###  Culture
- Cafe locations
- Libraries
- Parks
- Museums
- Historical landmarks

###  Apocalypse Readiness
- Nuclear bomb radius
- Rainfall data
- Flood data
- Elevation

###  Accessibility
- Bus stations
- Train stations

###  Commerce
- Unemployment rate

###  Safety
- Unemployment rate
- Taxation data
- WAPOL Incidents
- WAPOL Offences
