package database

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/model"
	"database/sql"
	"fmt"
	"log"
)

func GetReviews(db *sql.DB, listingId int, c chan []model.Review) {
	title := fmt.Sprintf("Get all reviews for listing id %v", listingId)
	query := `SELECT id, name, review FROM Review WHERE listing_id=?`
	result, err := db.Query(query, listingId)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- nil
		return
	}
	var review model.Review
	var reviews []model.Review
	for result.Next() {
		err = result.Scan(&review.Id, &review.Name, &review.Review)
		if err != nil {
			log.Printf("%v: %v\n", title, err)
			c <- nil
			return
		}
		reviews = append(reviews, review)
	}
	log.Printf("%v: successfully retrieve details of listing id %v\n", title, listingId)
	c <- reviews
}

func AddReview(db *sql.DB, newReview model.AddReview, c chan error) {
	title := fmt.Sprintf("Add new review for listing id of %v", newReview.Listing_id)
	_, err := db.Exec("INSERT INTO Review VALUES(?,?,?,?)", newReview.Id, newReview.Name, newReview.Review, newReview.Listing_id)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	log.Printf("%v: successfully added new review\n", title)
	c <- nil
}
