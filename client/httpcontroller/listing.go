package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Listing struct {
	Id              int
	ShopTitle       string
	ShopDescription string
	IgURL           string
	FbURL           string
	WebsiteURL      string
	BownerID        int
	CategoryID      int
}

func GetAllListing(c chan []Listing) {
	response, err := http.Get(BaseURL + "/listing")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		if response.StatusCode == http.StatusOK {
			defer response.Body.Close()
			data, _ := ioutil.ReadAll(response.Body)
			var listings map[string]Listing
			json.Unmarshal([]byte(data), &listings)
		}
	}
}

func ProcessAddListing(listing Listing, c chan error) {
	listingJSON, _ := json.Marshal(listing)
	response, err := http.Post(BaseURL+"/listing/add", "application/json", bytes.NewBuffer(listingJSON))
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

func ProcessDeleteListing(listing_id string, c chan error) {
	request, _ := http.NewRequest(http.MethodDelete, BaseURL+"/listing"+listing_id, nil)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		c <- err
		return
	}
	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		err := errors.New(string(data))
		c <- err
		return
	}
	c <- nil
}
