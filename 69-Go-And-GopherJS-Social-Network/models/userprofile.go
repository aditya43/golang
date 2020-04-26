package models

import "github.com/github.com/aditya43/golang-core/69-Go-And-GopherJS-Social-Network/forms"

type UserProfile struct {
	PageTitle        string `json:"pageTitle", bson:"pageTitle"`
	UUID             string `json:"uuid" bson:"uuid"`
	Username         string `json:"username" bson:"username"`
	About            string `json:"about" bson:"about"`
	Location         string `json:"location" bson:"location"`
	Interests        string `json:"interests" bson:"interests"`
	ProfileImagePath string `json:"profileImagePath" bson:"profileImagePath"`
	Form             *forms.MyProfileForm
}
