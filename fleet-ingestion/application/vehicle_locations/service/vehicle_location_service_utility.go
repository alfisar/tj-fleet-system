package service

import "github.com/umahmood/haversine"

func isRadius(lat1 float64, lon1 float64, lat2 float64, lon2 float64) bool {
	pointOne := haversine.Coord{
		Lat: lat1,
		Lon: lon1,
	}

	pointTwo := haversine.Coord{
		Lat: lat2,
		Lon: lon2,
	}

	_, km := haversine.Distance(pointOne, pointTwo)
	return km <= 0.05
}
