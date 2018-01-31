package utils

import (
	"fmt"

	geojson "github.com/paulmach/go.geojson"
)

func MarchallingToGeoJson(restaurants []Restaurant) (string, error) {
	fc := geojson.NewFeatureCollection()

	for _, restaurant := range restaurants {
		lat := restaurant.Lat
		lng := restaurant.Lng
		f := geojson.NewPointFeature([]float64{lat, lng})
		f.SetProperty("name", restaurant.Name)
		f.SetProperty("genre", restaurant.Genre)
		f.SetProperty("message", restaurant.Message)
		f.SetProperty("url", restaurant.URL)
		f.SetProperty("address", restaurant.Address)
		fc.AddFeature(f)
	}

	rawJSON, err := fc.MarshalJSON()

	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return "", err
	}

	return string(rawJSON), nil
}
