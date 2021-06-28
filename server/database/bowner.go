package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type BOwner struct {
	Id       int
	Email    string
	Password string
	Contact  string
}

func AddBOwner(db *sql.DB, email, password, contact string, c chan error) {
	encryptedpw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	_, err := db.Exec("INSERT INTO proj_db.BOwner VALUES(?,?,?,?)", 0, email, encryptedpw, contact)
	if err != nil {
		log.Println(err)
		c <- err
		return
	}
	c <- nil
}

func VerifyBOwnerIdentity(db *sql.DB, email string, password string, c chan error) {
	result, err := db.Query("SELECT * FROM sample_user.User WHERE email= ?", email)
	if err != nil {
		c <- err
		return
	}
	// get result
	var credentials BOwner
	for result.Next() {
		err = result.Scan(&credentials.Id, &credentials.Email, &credentials.Password, &credentials.Contact)
		if err != nil {
			c <- err
			return
		}
	}
	// check if email exists
	if credentials.Email == "" {
		c <- errors.New("no user found")
		fmt.Println("no user found")
		return
	}
	// check if password matches
	err = bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password))
	if err != nil {
		c <- errors.New("password do not match")
		fmt.Println("password do not match")
		return
	}
	c <- nil
}
