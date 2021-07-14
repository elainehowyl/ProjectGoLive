package database

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/model"
	"database/sql"
	"fmt"
	"log"
)

// type Listing struct {
// 	Id              int
// 	ShopTitle       string
// 	ShopDescription string
// 	IgURL           string
// 	FbURL           string
// 	WebsiteURL      string
// 	BOwner_id       int
// 	Category_id     int
// }

// func GetListings(db *sql.DB) {
// 	result, err := db.Query(
// 		`SELECT * FROM proj_db.Listing
// 		JOIN proj_db.BOwner
// 		ON proj_db.Listing.bowner_id = proj_db.BOwner.id
// 		JOIN proj_db.Category
// 		ON proj_db.Listing.category_id = proj_db.Category.id
// 	`)
// }

func GetOneListing(db *sql.DB, listingId int, c chan *model.Listing_Items) {
	title := fmt.Sprintf("Get details for listing id %v", listingId)
	query := `SELECT Listing.id, shop_title, shop_description, ig_url, fb_url, website_url, Category.id, title
		FROM Listing
		JOIN Category
		ON Listing.category_id=Category.id
		WHERE Listing.id=?`
	result, err := db.Query(query, listingId)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- nil
		return
	}
	var listing model.Listing_Items
	for result.Next() {
		err = result.Scan(&listing.Id, &listing.ShopTitle, &listing.ShopDescription, &listing.IgURL,
			&listing.FbURL, &listing.WebsiteURL, &listing.Category.Id, &listing.Category.Title)
		if err != nil {
			log.Printf("%v: %v\n", title, err)
			c <- nil
			return
		}
	}
	log.Printf("%v: successfully retrieve details of listing id %v\n", title, listingId)
	c <- &listing
}

func GetOneListingWReviews(db *sql.DB, listingId int, c chan *model.Listing_Items_Reviews) {
	title := fmt.Sprintf("Get details including reviews for listing id %v", listingId)
	query := `SELECT Listing.id, shop_title, shop_description, ig_url, fb_url, website_url, Category.id, title
		FROM Listing
		JOIN Category
		ON Listing.category_id=Category.id
		WHERE Listing.id=?`
	result, err := db.Query(query, listingId)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- nil
		return
	}
	var listing model.Listing_Items_Reviews
	for result.Next() {
		err = result.Scan(&listing.Id, &listing.ShopTitle, &listing.ShopDescription, &listing.IgURL,
			&listing.FbURL, &listing.WebsiteURL, &listing.Category.Id, &listing.Category.Title)
		if err != nil {
			log.Printf("%v: %v\n", title, err)
			c <- nil
			return
		}
	}
	log.Printf("%v: successfully retrieve details of listing id %v\n", title, listingId)
	c <- &listing
}

func GetListings(db *sql.DB, email string, c chan []model.Listing) {
	title := fmt.Sprintf("Get %v's listings", email)
	query := `SELECT Listing.id, shop_title, shop_description, ig_url, fb_url, website_url, Category.id, title
		FROM Listing
		JOIN Category
		ON Listing.category_id=Category.id
		WHERE bowner_email=?`
	result, err := db.Query(query, email)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		//c <- nil
		return
	}
	var listing model.Listing
	var listings []model.Listing
	for result.Next() {
		err = result.Scan(&listing.Id, &listing.ShopTitle, &listing.ShopDescription, &listing.IgURL,
			&listing.FbURL, &listing.WebsiteURL, &listing.Category.Id, &listing.Category.Title)
		if err != nil {
			log.Printf("%v: %v\n", title, err)
			//c <- nil
			return
		}
		listings = append(listings, listing)
	}
	if listings == nil {
		log.Printf("%v: no existing listings found\n", title)
	} else {
		log.Printf("%v: successfully retrieve %v's listing\n", title, email)
	}
	c <- listings
}

func AddListing(db *sql.DB, newListing model.Listing, email string, c chan error) {
	title := "Add New Listing"
	_, err := db.Exec("INSERT INTO Listing VALUES(?,?,?,?,?,?,?,?)", newListing.Id, newListing.ShopTitle, newListing.ShopDescription, newListing.IgURL, newListing.FbURL, newListing.WebsiteURL, email, newListing.Category.Id)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	log.Printf("%v: successfully added new listing into table\n", title)
	c <- nil
}

func DeleteListing(db *sql.DB, listingId int, c chan error) {
	title := "Delete Listing"
	_, err := db.Exec("DELETE FROM Listing WHERE id=?", listingId)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	log.Printf("%v: successfully deleted listing of id: %v\n", title, listingId)
	c <- nil
}

func GetAllListings(db *sql.DB, c chan []model.Listing) {
	title := "Get all listings in the table"
	query := `SELECT Listing.id, shop_title, shop_description, ig_url, fb_url, website_url, Category.id, title
		FROM Listing
		JOIN Category
		ON Listing.category_id=Category.id`
	result, err := db.Query(query)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		//c <- nil
		return
	}
	var listing model.Listing
	var listings []model.Listing
	for result.Next() {
		err = result.Scan(&listing.Id, &listing.ShopTitle, &listing.ShopDescription, &listing.IgURL,
			&listing.FbURL, &listing.WebsiteURL, &listing.Category.Id, &listing.Category.Title)
		if err != nil {
			log.Printf("%v: %v\n", title, err)
			//c <- nil
			return
		}
		listings = append(listings, listing)
	}
	if listings == nil {
		log.Printf("%v: no listings have been created yet\n", title)
	} else {
		log.Printf("%v: successfully retrieve all listings\n", title)
	}
	c <- listings
}
