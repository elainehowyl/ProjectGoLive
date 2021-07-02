package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func ProcessAddReview(review map[string]interface{}, c chan error) {
	reviewJSON, _ := json.Marshal(review)
	response, err := http.Post(BaseURL+"/review/add", "application/json", bytes.NewBuffer(reviewJSON))
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
