package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Customer struct {
	Id       int
	Email    string
	Username string
	Password string
}

func ProcessCustomerRegistration(customer Customer, c chan error) {
	customerJSON, _ := json.Marshal(customer)
	response, err := http.Post(BaseURL+"/customer/register", "application/json", bytes.NewBuffer(customerJSON))
	if err != nil {
		c <- err
		return
	}
	if response.StatusCode != http.StatusCreated {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		err = errors.New(string(data))
		c <- err
		return
	}
	c <- nil
}

func ProcessCustomerLogin(credentials map[string]string, c chan error) {
	credentialsJSON, _ := json.Marshal(credentials)
	response, err := http.Post(BaseURL+"/customer/login", "application/json", bytes.NewBuffer(credentialsJSON))
	if err != nil {
		c <- err
		return
	}
	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		err := errors.New(string(data))
		fmt.Println("err on processlogin function:", err)
		c <- err
		return
	}
	responseCookies := response.Cookies()
	for _, cookie := range responseCookies {
		if cookie.Name == "myCookie" {
			MyCookie = cookie
		}
	}
	c <- nil
}

func GetCustomer(loginCredentials map[string]string) {}
