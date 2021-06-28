package httpcontroller

import (
	"bytes"
	"encoding/json"
	"errors"
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
	response, err := http.Post("baseURL", "application/json", bytes.NewBuffer(bownerJSON))
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

func GetBOwner(loginCredentials map[string]string, c chan error) {
	loginCredentialsJSON, _ := json.Marshal(loginCredentials)
	response, err := http.Post("baseURL", "application/json", bytes.NewBuffer(loginCredentialsJSON))
	if err != nil {
		c <- err
		return
	}
}
