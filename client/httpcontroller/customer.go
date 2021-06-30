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
	Email    string
	Username string
	Password string
	Token    string
}

func ProcessCustomerRegistration(customer map[string]string, c chan error) {
	customerJSON, _ := json.Marshal(customer)
	response, err := http.Post(baseURL+"/customer/register", "application/json", bytes.NewBuffer(customerJSON))
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
	response, err := http.Post(baseURL+"/customer/login", "application/json", bytes.NewBuffer(credentialsJSON))
	if err != nil {
		c <- err
		return
	}
	//defer response.Body.Close()
	fmt.Println("response code:", response.StatusCode)
	//myHeader := response.Header
	//fmt.Println("RESPONSE HEADER:", myHeader.Get("myCookie"))
	responseCookie := response.Cookies()
	fmt.Println("GET COOKIE FROM SERVER:", responseCookie)
	if response.StatusCode != http.StatusOK {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		err := errors.New(string(data))
		fmt.Println("err on processlogin function:", err)
		c <- err
		return
	}
	c <- nil
}

func GetCustomer(loginCredentials map[string]string) {}
