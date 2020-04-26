package datasource

import (
	"github.com/bbcloudGroup/gothic/database"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	*gorm.DB
}

func NewAdmin() Admin {
	return Admin{database.G()}
}

type User struct {
	gorm.Model
	Username	string	`gorm:"size:100;unique_index;not null;"`
	Password 	string	`gorm:"size:100;not null;"`
	Name 	 	string	`gorm:"size:100;not null;"`
	Avatar		string  `gorm:"size:255"`
	RememberToken string `gorm:"size:100"`
}
