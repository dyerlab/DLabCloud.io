package models

import (
	"errors"
	"github.com/jinzhu/gorm"
)

var ErrNoRecord = errors.New("models: no matching record found")



// Manuscript is the base type for a journal article
type Manuscript struct {
	gorm.Model
	Authors  string `json:"authors"`
	Year     int    `json:"year"`
	Title    string `json:"title"`
	Journal  string `json:"journal"`
	Volume   int    `json:"volume"`
	Pages    string `json:"pages"`
	Doi      string `json:"doi"`
	Abstract string `json:"abstract"`
}

// Snippet is a short piece of code to do some analysis
type Snippet struct {
	gorm.Model
	Title 		string `gorm:"column:name" json:"title"`
	Contents 	string `gorm:"colum:contents" json:"contents"`
}

