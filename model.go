package blogapi

import "github.com/jinzhu/gorm"

type (

	// Blog Db Structure
	Blog struct {
		gorm.Model
		Name     string `gorm:"not null;unique"`
		Comments []Comment
	}

	// Comment Db Structure
	Comment struct {
		gorm.Model
		Text   string `gorm:"not null;"`
		Blog   Blog
		BlogID int
	}
)
