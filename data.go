package main

import (
	_ "github.com/lib/pq"
	"log"
    "database/sql"
	"os"
//	http://l-lin.github.io/2015/01/31/Golang-Deploy_to_heroku
//	https://godoc.org/github.com/lib/pq
)

// // Fetch the list of novels
// func GetList() []*Novel {
// 	novels := make([]*Novel, 0)
// 	database := connect()
// 	defer database.Close()

// 	rows, err := database.Query("SELECT id, title, url, image_url, summary, favorite FROM novels")
// 	if err != nil {
// 		log.Fatalf("[x] Error when getting the list of novels. Reason: %s", err.Error())
// 	}
// 	for rows.Next() {
// 		n := toNovel(rows)
// 		if n.IsValid() {
// 			novels = append(novels, n)
// 		}
// 	}
// 	if err := rows.Err(); err != nil {
// 		log.Fatalf("[x] Error when getting the list of novels. Reason: %s", err.Error())
// 	}
// 	return novels
// }

// Connect to Heroku database using the OS env DATABASE_URL
func connect() *sql.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	database, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("[x] Could not open the connection to the database. Reason: %s", err.Error())
	}
	return database
}

// // Fetch the content of the rows and build a new novel
// func toNovel(rows db.RowMapper) *Novel {
// 	var id string
// 	var title string
// 	var url string
// 	var imageUrl string
// 	var summary string
// 	var favorite bool

// 	rows.Scan(&id, &title, &url, &imageUrl, &summary, &favorite)

// 	return &Novel{
// 		Id: id,
// 		Title: title,
// 		Url: url,
// 		ImageUrl: imageUrl,
// 		Summary: summary,
// 		Favorite: favorite,
// 	}
// }