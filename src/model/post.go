package post

import (
	"database/sql"
	"errors"
	"fmt"
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

type GetRowResponse struct {
	PostDisplay PostDisplay
	Valid       bool
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

func GetRow(id string) GetRowResponse {
	db := db.Execute()
	defer db.Close()
	post := &Post{}
	display := &PostDisplay{}
	res := &GetRowResponse{}
	err := db.QueryRow("SELECT * FROM posts WHERE id = ?", id).
		Scan(&post.ID, &post.Title, &post.Content, &post.Created)
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("GetRow no records.")
		res.PostDisplay = *display
		res.Valid = false
		return *res
	}
	if err != nil {
		log.Fatalf("GetRow db.QueryRow error err:%v", err)
	}
	display.ID = post.ID
	display.Title = post.Title
	display.Created = post.Created.Format("2006/1/2 15:04")
	if post.Content.Valid {
		display.Content = post.Content.String
	} else {
		display.Content = ""
	}
	res.PostDisplay = *display
	res.Valid = true
	return *res
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
