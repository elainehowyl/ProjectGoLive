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
	response, err := http.Post(baseURL+"/item/add", "application/json", bytes.NewBuffer(itemJSON))
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

func ProcessUpdateItem(item_id string, item map[string]interface{}, c chan error) {
	itemJSON, _ := json.Marshal(item)
	request, _ := http.NewRequest(http.MethodPut, baseURL+"/item/"+item_id, bytes.NewBuffer(itemJSON))
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

func ProcessDeleteItem(item_id string, c chan error) {
	request, _ := http.NewRequest(http.MethodDelete, baseURL+"/item/"+item_id, nil)
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
