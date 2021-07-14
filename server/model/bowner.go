package model

type BOwner struct {
	Email    string
	Password string
	Contact  string
}

type BOwnerCredentials struct {
	Email    string
	Password string
}

type BOwnerData struct {
	Email    string
	Contact  string
	Listings []Listing
}

// type BOwnerDetails struct {
// 	Id      int
// 	Email   string
// 	Contact string
// 	Listing *Listing
// }
