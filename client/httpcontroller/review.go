package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ProcessAddReview(review map[string]string, c chan error) {
	reviewJSON, _ := json.Marshal(review)
	response, err := http.Post(baseURL+"/review/add", "application/json", bytes.NewBuffer(reviewJSON))
	if err != nil {
		c <- err
		return
	}
	if response.StatusCode != http.StatusCreated {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		err := errors.New(string(data))
		fmt.Println("err on processlogin function:", err)
		c <- err
		return
	}
}
