package models

import (
	"database/sql"
	_ "log"
)
var db *sql.DB

type Post struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Body    string `json:"body"`
}

func GetPost () Post {
    var post Post
    return post
}

func GetPosts() []Post {
    var posts []Post
    return posts
}





