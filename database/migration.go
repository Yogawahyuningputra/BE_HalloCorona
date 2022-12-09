package database

import (
	"be_corona/models"
	"be_corona/pkg/mysql"
	"fmt"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Article{}, models.Consultation{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
