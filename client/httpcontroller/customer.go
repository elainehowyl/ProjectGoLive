package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Customer struct {
	Email    string
	Username string
	Password string
}

func AddCustomer(customer map[string]string) error {
	customerJSON, _ := json.Marshal(customer)
	response, err := http.Post("baseURL", "application/json", bytes.NewBuffer(customerJSON))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusCreated {
		data, _ := ioutil.ReadAll(response.Body)
		err = errors.New(string(data))
		return err
	}
	return nil
}
