package database

import (
	"database/sql"
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
	encryptedpw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	_, err := db.Exec("INSERT INTO sample_user.User VALUES(?,?,?)", 0, email, username, encryptedpw)
	if err != nil {
		log.Println(err)
		c <- err
		return
	}
	c <- nil
}
