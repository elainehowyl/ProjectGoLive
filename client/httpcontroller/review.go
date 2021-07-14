package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

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

func ProcessAddReview(review AddReview, c chan error) {
	reviewJSON, _ := json.Marshal(review)
	response, err := http.Post(BaseURL+"/listing/"+strconv.Itoa(review.Listing_id)+"/review/add", "application/json", bytes.NewBuffer(reviewJSON))
	if err != nil {
		c <- err
		return
	}
	if response.StatusCode != http.StatusCreated {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		err := errors.New(string(data))
		c <- err
		return
	}
	c <- nil
}
