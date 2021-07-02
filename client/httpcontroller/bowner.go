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
	ID       int
	Email    string
	Contact  string
	Password string
}

func ProcessBOwnerRegistration(bowner BOwner, c chan error) {
	bownerJSON, _ := json.Marshal(bowner)
	response, err := http.Post(BaseURL+"/bowner/register", "application/json", bytes.NewBuffer(bownerJSON))
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
