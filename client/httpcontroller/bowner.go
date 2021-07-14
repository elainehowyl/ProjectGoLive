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
}

// type BOwnerDetails struct {
// 	ID      int
// 	Email   string
// 	Contact string
// }

type Profile struct {
	Email    string
	Contact  string
	Listings []Listing
}

var CurrentUser *Profile

func ProcessBOwnerRegistration(bowner BOwner, c chan error) {
	bownerJSON, _ := json.Marshal(bowner)
	response, err := http.Post(BaseURL+"/register", "application/json", bytes.NewBuffer(bownerJSON))
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
	response, err := http.Post(BaseURL+"/login", "application/json", bytes.NewBuffer(credentialsJSON))
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

func GetBOwnerData(email string, c chan *Profile) {
	req, err := http.NewRequest(http.MethodGet, BaseURL+"/profile/"+email, nil)
	if err != nil {
		fmt.Println(err.Error())
		c <- nil
		return
	}
	req.AddCookie(MyCookie)
	response, err2 := Client.Do(req)
	if err2 != nil {
		fmt.Println(err2.Error())
		c <- nil
		return
	}
	responseCookies := response.Cookies()
	for _, cookie := range responseCookies {
		if cookie.Name == "myCookie" {
			MyCookie = cookie
		}
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		fmt.Println(string(data))
		c <- nil
		return
	}
	// var profile *Profile
	// json.Unmarshal([]byte(data), &profile)
	// c <- profile
	json.Unmarshal([]byte(data), &CurrentUser)
	c <- CurrentUser
}

func ProcessBOwnerLogout(email string, c chan error) {
	emailJSON, _ := json.Marshal(email)
	req, err := http.NewRequest(http.MethodPost, BaseURL+"/logout", bytes.NewBuffer(emailJSON))
	if err != nil {
		fmt.Println(err.Error())
		c <- err
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.AddCookie(MyCookie)
	response, err2 := Client.Do(req)
	if err2 != nil {
		fmt.Println(err2.Error())
		c <- err
		return
	}
	if response.StatusCode != http.StatusNoContent {
		defer response.Body.Close()
		data, _ := ioutil.ReadAll(response.Body)
		err3 := errors.New(string(data))
		c <- err3
		return
	}
	c <- nil
}
