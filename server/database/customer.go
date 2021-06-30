package database

import (
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	Id       int
	Email    string
	Username string
	Password string
}

func AddCustomer(db *sql.DB, email, password, username string, c chan error) {
	title := "Register"
	encryptedpw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	_, err := db.Exec("INSERT INTO proj_db.Customer VALUES(?,?,?,?)", 0, email, username, encryptedpw)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	log.Printf("%v: successfully added new customer into table\n", title)
	c <- nil
}

func VerifyCustomerIdentity(db *sql.DB, email string, password string, c chan error) {
	title := "Login"
	result, err := db.Query("SELECT * FROM proj_db.Customer WHERE email= ?", email)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	// get result
	var credentials Customer
	for result.Next() {
		err = result.Scan(&credentials.Id, &credentials.Email, &credentials.Username, &credentials.Password)
		if err != nil {
			log.Printf("%v: %v\n", title, err)
			c <- err
			return
		}
	}
	// check if email exists
	if credentials.Email == "" {
		err = errors.New("no user found")
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	// check if password matches
	err = bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password))
	if err != nil {
		err = errors.New("password do not match")
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	log.Printf("%v: customer verification passed", title)
	c <- nil
}
