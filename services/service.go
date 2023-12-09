package services

import (
    "encoding/json"
    "log"
    "database/sql"
    "net/http"
    "github.com/jmoiron/sqlx"
    "github.com/tbilous/go_server/models"
)

var db_conn *sqlx.DB

func GetAllPosts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var posts = models.GetPosts()

	sqlStmt := `SELECT * FROM posts`
	rows, err := db_conn.Queryx(sqlStmt)

	if err == nil {
		var tempPost = models.GetPost()

		for rows.Next() {
			err = rows.StructScan(&tempPost)
			posts = append(posts, tempPost)
		}

		switch err {
		case sql.ErrNoRows:
			{
				log.Println("No rows returned.")
				http.Error(w, err.Error(), 204)
			}
		case nil:
			json.NewEncoder(w).Encode(&posts)
		default:
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		http.Error(w, err.Error(), 400)
		return
	}
}

func SetDB(db *sqlx.DB) {
    db_conn = db
}
