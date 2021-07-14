package model

type Review struct {
	Id     int
	Name   string
	Review string
	//Listing_id int
}

type AddReview struct {
	Id         int
	Name       string
	Review     string
	Listing_id int
}
