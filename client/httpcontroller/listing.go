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
