package pkg

import "strconv"

//
func GetGeolocationURL(latitude float64, longitude float64) string {
	coordinate := "https://maps.google.com/maps?q=" + strconv.FormatFloat(latitude, 'f', -1, 64) + "," + strconv.FormatFloat(longitude, 'f', -1, 64)
	return coordinate
}
