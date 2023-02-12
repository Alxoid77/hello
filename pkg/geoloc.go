package pkg

import (
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func GetGeolocationURL(chatid int64, user string, latitude float64, longitude float64) string {
	coordinate := "https://maps.google.com/maps?q=" + strconv.FormatFloat(latitude, 'f', -1, 64) + "," + strconv.FormatFloat(longitude, 'f', -1, 64)
	/*
		db, err := sql.Open("sqlite3", "./db.db")
		if err != nil {
			panic(err)
		}
		defer db.Close()
		result, err := db.Prepare("insert into users (ChatID, UserName, Latitude, Longitude) values (5551212, 'vasyanil', latitude, longitude)")

		if err != nil {
			panic(err)
		} else {
			result.Exec()
		}
	*/
	return coordinate
}
