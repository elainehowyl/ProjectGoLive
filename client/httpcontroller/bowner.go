package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BOwner struct {
	Email    string
	Contact  string
	Password string
	Token    string
}

func AddBOwner(bowner map[string]string, c chan error) {
	bownerJSON, _ := json.Marshal(bowner)
	response, err := http.Post("http://localhost:5000/bowner/register", "application/json", bytes.NewBuffer(bownerJSON))
	if err != nil {
		c <- err
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusCreated {
		data, _ := ioutil.ReadAll(response.Body)
		err := errors.New(string(data))
		c <- err
		return
	}
	c <- nil
}

func ProcessBOwnerLogin(credentials map[string]string, c chan error) {
	credentialsJSON, _ := json.Marshal(credentials)
	response, err := http.Post("http://localhost:5000/customer/login", "application/json", bytes.NewBuffer(credentialsJSON))
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

// func GetBOwner(loginCredentials map[string]string, c chan error) {
// 	loginCredentialsJSON, _ := json.Marshal(loginCredentials)
// 	response, err := http.Post("baseURL", "application/json", bytes.NewBuffer(loginCredentialsJSON))
// 	if err != nil {
// 		c <- err
// 		return
// 	}
// }
