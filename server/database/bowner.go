package database

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/model"
	"database/sql"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// type BOwner struct {
// 	Id       int
// 	Email    string
// 	Password string
// 	Contact  string
// }

// type BOwnerCredentials struct {
// 	Email    string
// 	Password string
// }

// type BOwnerDetails struct {
// 	Id      int
// 	Email   string
// 	Contact string
// }

func AddBOwner(db *sql.DB, email, password, contact string, c chan error) {
	title := "Register"
	encryptedpw, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	_, err := db.Exec("INSERT INTO proj_db.BOwner VALUES(?,?,?)", email, encryptedpw, contact)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	log.Printf("%v: successfully added new business owner into table\n", title)
	c <- nil
}

func VerifyBOwnerIdentity(db *sql.DB, email string, password string, c chan error) {
	title := "Login"
	result, err := db.Query("SELECT * FROM proj_db.BOwner WHERE email= ?", email)
	//result, err := db.Query("SELECT id, email, contact WHERE proj_db.BOwner WHERE email=?", email)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- err
		return
	}
	// get result
	var credentials model.BOwner
	for result.Next() {
		err = result.Scan(&credentials.Email, &credentials.Password, &credentials.Contact)
		if err != nil {
			log.Printf("%v: %v\n", title, err)
			c <- err
			return
		}
	}
	// check if email exists
	if credentials.Email == "" {
		err = errors.New("no user found")
		c <- err
		log.Printf("%v: %v\n", title, err)
		return
	}
	// check if password matches
	err = bcrypt.CompareHashAndPassword([]byte(credentials.Password), []byte(password))
	if err != nil {
		err = errors.New("password do not match")
		c <- err
		log.Printf("%v: %v\n", title, err)
		return
	}
	log.Printf("%v: business owner verification passed\n", title)
	c <- nil
}

func GetBOwnerData(db *sql.DB, email string, c chan *model.BOwnerData) {
	title := "Retrieve Business Owner Info"
	// query := `SELECT BOwner.id, email, contact, Listing.id, shop_title, shop_description, ig_url, fb_url, website_url, Category.id, title
	// FROM BOwner
	// JOIN Listing ON BOwner.id=Listing.bowner_id
	// JOIN Category ON Category.id=Listing.category_id
	// WHERE email=?`
	query := `SELECT email, contact FROM BOwner WHERE email=?`
	result, err := db.Query(query, email)
	if err != nil {
		log.Printf("%v: %v\n", title, err)
		c <- nil
		return
	}
	//var details model.BOwnerDetails
	var data model.BOwnerData
	for result.Next() {
		err = result.Scan(&data.Email, &data.Contact)
		if err != nil {
			log.Printf("%v: %v\n", title, err)
			c <- nil
			return
		}
	}
	log.Printf("%v: successfully retrieved business owner's details", title)
	c <- &data
}
