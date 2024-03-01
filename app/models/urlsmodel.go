package models

type URLS struct {
	ShortPath   string `json:"shortpath" bson:"shortpath"`
	Destination string `json:"destination" bson:"destination"`
}