package post

import (
	"database/sql"
	"log"
	"time"

	db "github.com/Tomoya185-miyawaki/go-blog/utils"
)

type Post struct {
	ID      int
	Title   string
	Content sql.NullString
	Created time.Time
}

type PostDisplay struct {
	ID      int
	Title   string
	Content string
	Created string
}

func GetRows() []PostDisplay {
	db := db.Execute()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatalf("GetRows db.Query error err:%v", err)
	}
	defer rows.Close()

	posts := []PostDisplay{}
	for rows.Next() {
		p := Post{}
		display := PostDisplay{}
		if err := rows.Scan(&p.ID, &p.Title, &p.Content, &p.Created); err != nil {
			log.Fatalf("GetRows rows.Scan error err:%v", err)
		}
		display.ID = p.ID
		display.Title = p.Title
		display.Created = p.Created.Format("2006/1/2 15:04")
		if p.Content.Valid {
			display.Content = p.Content.String
		} else {
			display.Content = ""
		}
		posts = append(posts, display)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("GetRows rows.Err error err:%v", err)
	}
	return posts
}

func Create(title string, content string) {
	db := db.Execute()
	defer db.Close()

	insert, err := db.Prepare("INSERT INTO posts(title, content) VALUES(?, ?)")
	if err != nil {
		log.Fatalf("Create err:%v", err)
	}
	_, err = insert.Exec(title, content)
	if err != nil {
		panic(err.Error())
	}
}
