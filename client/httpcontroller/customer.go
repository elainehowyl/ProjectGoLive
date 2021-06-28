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
	Token    string
}

func AddCustomer(customer map[string]string) error {
	customerJSON, _ := json.Marshal(customer)
	response, err := http.Post("baseURL", "application/json", bytes.NewBuffer(customerJSON))
	if err != nil {
		// c <- err
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusCreated {
		data, _ := ioutil.ReadAll(response.Body)
		err = errors.New(string(data))
		// c <- err
		return err
	}
	//c <- nil
	return nil
}

func GetCustomer(loginCredentials map[string]string) {}
