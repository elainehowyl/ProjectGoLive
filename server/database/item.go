package database

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/model"
	"database/sql"
	"fmt"
	"log"
)

func GetItems(db *sql.DB, listingId int, c chan []model.Item) {
	title := fmt.Sprintf("Get Items with listing id of %v", listingId)
	result, err := db.Query("SELECT * FROM Item WHERE listing_id=?", listingId)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- nil
		return
	}
	var item model.Item
	var items []model.Item
	for result.Next() {
		err = result.Scan(&item.Id, &item.Name, &item.Price, &item.Description, &item.ListingId)
		if err != nil {
			log.Printf("%v: %v\n", title, err)
			c <- nil
			return
		}
		item.Price = float64(item.Price / 100)
		items = append(items, item)
	}
	if items == nil {
		log.Printf("%v: no existing items found for listing id of %v\n", title, listingId)
	} else {
		log.Printf("%v: successfully retrieve items for listing id of %v\n", title, listingId)
	}
	c <- items
}

func AddItem(db *sql.DB, newItem model.Item, c chan error) {
	title := "Add New Item"
	_, err := db.Exec("INSERT INTO Item VALUES(?,?,?,?,?)", newItem.Id, newItem.Name, int(newItem.Price*100), newItem.Description, newItem.ListingId)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	log.Printf("%v: successfully added new item for listing id of %v\n", title, newItem.ListingId)
	c <- nil
}

func DeleteItem(db *sql.DB, listingId int, c chan error) {
	title := fmt.Sprintf("Delete all items associate with listing id of %v", listingId)
	_, err := db.Exec("DELETE FROM Item WHERE listing_id=?", listingId)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	log.Printf("%v: successfully deleted all items with listing id of %v\n", title, listingId)
	c <- nil
}

func DeleteSingleItem(db *sql.DB, itemId int, c chan error) {
	title := "Delete Single Item"
	_, err := db.Exec("DELETE FROM Item WHERE id=?", itemId)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	log.Printf("%v: successfully deleted item with id of %v\n", title, itemId)
	c <- nil
}
