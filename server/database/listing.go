package database

import (
	"database/sql"
	"log"
)

type Listing struct {
	Id              int
	ShopTitle       string
	ShopDescription string
	IgURL           string
	FbURL           string
	WebsiteURL      string
	BOwner_id       int
	Category_id     int
}

// func GetListings(db *sql.DB) {
// 	result, err := db.Query(
// 		`SELECT * FROM proj_db.Listing
// 		JOIN proj_db.BOwner
// 		ON proj_db.Listing.bowner_id = proj_db.BOwner.id
// 		JOIN proj_db.Category
// 		ON proj_db.Listing.category_id = proj_db.Category.id
// 	`)
// }

func AddListing(db *sql.DB, newListing Listing, c chan error) {
	title := "Add New Listing"
	_, err := db.Exec("INSERT INTO proj_db.Listing VALUES(?,?,?,?,?,?,?,?)", newListing.Id, newListing.ShopTitle, newListing.ShopDescription, newListing.IgURL, newListing.FbURL, newListing.WebsiteURL, newListing.BOwner_id, newListing.Category_id)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	log.Printf("%v: successfully added new listing into table\n", title)
	c <- nil
}
