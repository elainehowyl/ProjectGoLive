package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func ProcessAddItem(item map[string]interface{}, c chan error) {
	itemJSON, _ := json.Marshal(item)
	response, err := http.Post(baseURL+"/item/add", "application/json", bytes.NewBuffer([]byte(itemJSON)))
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
