package models

type User struct {
	ID 			int 		`gorm:"primaryKey"`
	Username 	string 		`gorm:"type:varchar(255)"`
	Password 	string 		`gorm:"type:varchar(255)"`
	Email 		string 		`gorm:"type:varchar(255)"`
	Name 		string 		`gorm:"type:varchar(255)"`
	Experience 	[]string  	`gorm:"type:varchar(255)[]"`
	Year		[]string	`gorm:"type:varchar(255)[]"`
}
