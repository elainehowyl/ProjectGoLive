package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Category struct {
	Id    int
	Title string
}

// type Listing struct {
// 	Id              int
// 	ShopTitle       string
// 	ShopDescription string
// 	IgURL           string
// 	FbURL           string
// 	WebsiteURL      string
// 	BownerID        int
// 	CategoryID      int
// }

type Listing struct {
	Id              int
	ShopTitle       string
	ShopDescription string
	IgURL           string
	FbURL           string
	WebsiteURL      string
	Category        Category
}

type Listing_Items struct {
	Id              int
	ShopTitle       string
	ShopDescription string
	IgURL           string
	FbURL           string
	WebsiteURL      string
	Category        Category
	Items           []Item
}

type Listing_Items_Reviews struct {
	Id              int
	ShopTitle       string
	ShopDescription string
	IgURL           string
	FbURL           string
	WebsiteURL      string
	Category        Category
	Items           []Item
	Reviews         []Review
}

func GetOneListing(listingId int, c chan *Listing_Items) {
	response, err := http.Get(BaseURL + "/" + CurrentUser.Email + "/listing/" + strconv.Itoa(listingId) + "/view")
	if err != nil {
		fmt.Println(err.Error())
		c <- nil
		return
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		fmt.Println(string(data))
		c <- nil
		return
	}
	var listingWItems *Listing_Items
	json.Unmarshal([]byte(data), &listingWItems)
	c <- listingWItems
}

func GetAllListings(c chan []Listing) {
	response, err := http.Get(BaseURL + "/listings")
	if err != nil {
		fmt.Println(err.Error())
		//c <- nil
		return
	}
	defer response.Body.Close()
	if response.StatusCode == http.StatusOK {
		data, _ := ioutil.ReadAll(response.Body)
		var listings []Listing
		json.Unmarshal([]byte(data), &listings)
		c <- listings
	}
}

func ProcessAddListing(listing Listing, c chan error) {
	listingJSON, _ := json.Marshal(listing)
	req, err := http.NewRequest(http.MethodPost, BaseURL+"/"+CurrentUser.Email+"/listing/add", bytes.NewBuffer(listingJSON))
	if err != nil {
		fmt.Println(err.Error())
		c <- nil
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(MyCookie)
	response, err2 := Client.Do(req)
	if err2 != nil {
		fmt.Println(err2.Error())
		c <- nil
		return
	}
	responseCookies := response.Cookies()
	for _, cookie := range responseCookies {
		if cookie.Name == "myCookie" {
			MyCookie = cookie
		}
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
	req, err := http.NewRequest(http.MethodDelete, BaseURL+"/"+CurrentUser.Email+"/listing/"+listing_id+"/delete", nil)
	if err != nil {
		fmt.Println(err.Error())
		c <- err
		return
	}
	req.AddCookie(MyCookie)
	response, err2 := Client.Do(req)
	if err2 != nil {
		c <- err
		return
	}
	responseCookies := response.Cookies()
	for _, cookie := range responseCookies {
		if cookie.Name == "myCookie" {
			MyCookie = cookie
		}
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

func GetListingWReviews(listingId int, c chan Listing_Items_Reviews) {
	response, err := http.Get(BaseURL + "/listing/" + strconv.Itoa(listingId))
	if err != nil {
		fmt.Println(err.Error())
		//c <- err
		return
	}
	if response.StatusCode == http.StatusOK {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		var listingWReviews Listing_Items_Reviews
		json.Unmarshal([]byte(data), &listingWReviews)
		c <- listingWReviews
	}
}
