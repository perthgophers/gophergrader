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

## Datasets

- WAPOL CAD Incidents
- WAPOL Offences
- [2074.0 - Census of Population and Housing: Mesh Block Counts, 2011](http://abs.gov.au/ausstats/abs@.nsf/Lookup/2074.0main+features12011)
- [FPM 100 Year ARI Floodplain Development Control Areas](http://atlases.water.wa.gov.au/idelve/dowdataext/metadata_statements/FPM_FDCA.html)
- [Medium Scale Points of Interest (LGATE-135)](http://catalogue.beta.data.wa.gov.au/dataset/medium-scale-points-of-interest-lgate-135)
- [Network Operations Traffic Data MRWA](http://catalogue.beta.data.wa.gov.au/dataset/network-operations-traffic-data-mrwa)
- [My Community Directory](https://directorycdn.blob.core.windows.net/govhack/Service%20Data%20GovHack%20WA%202016.xlsx)
- [Temperature and Rainfall Historical](https://data.gov.au/dataset/weather-forecasting-verification-data-2015-05-to-2016-04/resource/16083945-3309-4c8a-9b64-49971be15878)
- [WA Meshblock Metadata](http://abs.gov.au/ausstats/subscriber.nsf/log?openagent&1270055001_mb_2011_wa_shape.zip&1270.0.55.001&Data%20Cubes&E1F9D7EA5D7FE609CA257801000CD692&0&July%202011&23.12.2010&Latest)

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
