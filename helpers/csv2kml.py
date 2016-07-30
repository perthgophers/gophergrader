# From http://gis.stackexchange.com/questions/19163/how-to-convert-shp-to-csv-including-attributes-and-geometry

import ogr,csv,sys,os
ogr.UseExceptions()

csvfile=r'C:\temp\test.csv' #sys.argv[1]
kmlfile=r'C:\temp\test.kml' #sys.argv[2]

csvreader=csv.reader(open(csvfile,'rb'))
headers=csvreader.next()

ds = ogr.GetDriverByName('KML').CreateDataSource(kmlfile)
lyr = ds.CreateLayer(os.path.splitext(os.path.basename(kmlfile))[0])

for field in headers[:-1]: #skip kmlgeometry
    field_def = ogr.FieldDefn(field)
    print lyr.CreateField(field_def)

for rec in csvreader:
    feat = ogr.Feature(lyr.GetLayerDefn())
    for i,field in enumerate(headers[:-1]): #skip kmlgeometry
        feat.SetField(field, rec[i])
    feat.SetGeometry(ogr.CreateGeometryFromGML(rec[-1]))
    lyr.CreateFeature(feat)

del lyr,ds