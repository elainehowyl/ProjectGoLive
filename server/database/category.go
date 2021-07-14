package database

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/model"
	"database/sql"
	"log"
)

func GetCategory(db *sql.DB, CategoryTitle string, c chan int) {
	title := "Get Category"
	result, err := db.Query("SELECT id FROM Category WHERE title=?", CategoryTitle)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- 0
		return
	}
	var id int
	for result.Next() {
		err = result.Scan(&id)
		if err != nil {
			log.Printf("%v: %v\n", title, err)
			c <- 0
			return
		}
	}
	log.Printf("%v: successfully retrieved category", title)
	c <- id
}

func AddCategory(db *sql.DB, category model.Category, c chan int) {
	title := "Add Category"
	result, err := db.Exec("INSERT INTO Category VALUES(?,?)", category.Id, category.Title)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- 0
		return
	}
	id, _ := result.LastInsertId()
	log.Printf("%v: successfully added new category into table\n", title)
	c <- int(id)
}
