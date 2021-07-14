package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	Id          int
	Name        string
	Price       float64
	Description string
	ListingId   int
}

func ProcessAddItem(item Item, c chan error) {
	//"/{email}/listing/{listing_id}/item/add"
	// int to string
	listingId := strconv.Itoa(item.ListingId)
	itemJSON, _ := json.Marshal(item)
	req, err := http.NewRequest(http.MethodPost, BaseURL+"/"+CurrentUser.Email+"/listing/"+listingId+"/item/add", bytes.NewBuffer(itemJSON))
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
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode == http.StatusUnauthorized {
		err = errors.New(string(data))
		log.Println(err)
		c <- errors.New("unauthorized")
		return
	}
	responseCookies := response.Cookies()
	for _, cookie := range responseCookies {
		if cookie.Name == "myCookie" {
			MyCookie = cookie
		}
	}
	if response.StatusCode != http.StatusCreated {
		err = errors.New(string(data))
		c <- err
		return
	}
	c <- nil
}

func ProcessUpdateItem(item_id string, item map[string]interface{}, c chan error) {
	itemJSON, _ := json.Marshal(item)
	request, _ := http.NewRequest(http.MethodPut, BaseURL+"/item/"+item_id, bytes.NewBuffer(itemJSON))
	request.Header.Set("Content-Type", "application/json")
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

func ProcessDeleteItem(item_id, listing_id string, c chan error) {
	req, err := http.NewRequest(http.MethodDelete, BaseURL+"/"+CurrentUser.Email+"/listing/"+listing_id+"/item/"+item_id+"/delete", nil)
	if err != nil {
		//fmt.Println(err.Error())
		c <- err
		return
	}
	req.AddCookie(MyCookie)
	response, err2 := Client.Do(req)
	if err2 != nil {
		//log.Println(err)
		c <- err
		return
	}
	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		err := errors.New(string(data))
		//log.Println(err)
		c <- err
		return
	}
	c <- nil
}
