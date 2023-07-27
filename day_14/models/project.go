package models

import (
	"time"
)

type Project struct {
	// Name 		string		
	// StartDate 	time.Time
	// EndDate 	time.Time
	// Duration 	string
	// Description string
	// NodeJs     	bool
	// ReactJs    	bool
	// Golang     	bool
	// Java 		bool
	// Image 		string
	// Buat beserta jsonnya untuk database PostgreSql
	ID 			int 		`gorm:"primaryKey"`
	Name 		string 		`gorm:"type:varchar(255)"`
	StartDate 	time.Time 	`gorm:"type:date"`
	EndDate 	time.Time 	`gorm:"type:date"`
	Duration 	string 		`gorm:"type:varchar(255)"`
	Description string 		`gorm:"type:text"`
	NodeJs     	bool 		`gorm:"type:boolean"`
	ReactJs    	bool 		`gorm:"type:boolean"`
	Golang     	bool 		`gorm:"type:boolean"`
	Java 		bool 		`gorm:"type:boolean"`
	Image 		string 		`gorm:"type:varchar(255)"`
}

