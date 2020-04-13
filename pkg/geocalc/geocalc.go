package geocalc

import (
	"math"

	"github.com/fare-estimate/pkg/models"
)

const (
	earthDiameter = 6371000 // meters
	radPerDegree  = math.Pi / 180
)

func DistBetween(p1, p2 *models.PointInTime) float64 {
	lat1Rad := rad(p1.Lat)
	lat2Rad := rad(p2.Lat)

	deltaLat := rad(p2.Lat - p1.Lat)
	deltaLng := rad(p2.Lng - p1.Lng)

	msdlat := math.Sin(deltaLat / 2)
	msdlng := math.Sin(deltaLng / 2)

	a := msdlat*msdlat + math.Cos(lat1Rad)*math.Cos(lat2Rad)*msdlng*msdlng
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthDiameter * c
}

func rad(deg float64) float64 {
	return deg * radPerDegree
}
