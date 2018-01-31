package utils

import (
	"fmt"
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

func MarchallingToGeoJson(tabelogResult TabelogResult, geoCoordResult GeocoordResult) (string, error) {
	fc := geojson.NewFeatureCollection()

	lat, _ := strconv.ParseFloat(geoCoordResult.Coordinate.Lat.Text, 64)
	lng, _ := strconv.ParseFloat(geoCoordResult.Coordinate.Lng.Text, 64)
	f := geojson.NewPointFeature([]float64{lat, lng})
	f.SetProperty("name", tabelogResult.Name)
	f.SetProperty("url", tabelogResult.URL)
	f.SetProperty("address", tabelogResult.Address)
	fc.AddFeature(f)

	rawJSON, err := fc.MarshalJSON()

	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return "", err
	}

	return string(rawJSON), nil
}
