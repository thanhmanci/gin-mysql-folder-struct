package models

import "github.com/jinzhu/gorm"

type (
	// NewfeedModel describes a newfeedModel type
	NewfeedModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}
	// TransformedNewfeed represents a formatted newfeed
	TransformedNewfeed struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)
