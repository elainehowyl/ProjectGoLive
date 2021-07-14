package model

type Listing struct {
	Id              int
	ShopTitle       string
	ShopDescription string
	IgURL           string
	FbURL           string
	WebsiteURL      string
	//BOwner_id       int
	Category Category
}

type Listing_Items struct {
	Id              int
	ShopTitle       string
	ShopDescription string
	IgURL           string
	FbURL           string
	WebsiteURL      string
	Category        Category
	Items           []Item
}

type Listing_Items_Reviews struct {
	Id              int
	ShopTitle       string
	ShopDescription string
	IgURL           string
	FbURL           string
	WebsiteURL      string
	Category        Category
	Items           []Item
	Reviews         []Review
}
