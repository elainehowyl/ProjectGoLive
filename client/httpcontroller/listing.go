package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func ProcessAddListing(listing map[string]interface{}, c chan error) {
	listingJSON, _ := json.Marshal(listing)
	response, err := http.Post(baseURL+"/listing/add", "application/json", bytes.NewBuffer([]byte(listingJSON)))
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
	request, _ := http.NewRequest(http.MethodDelete, baseURL+"/listing"+listing_id, nil)
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
