package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:lx89031508989@tcp(localhost:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	//迁移schema
	// db.AutoMigrate(&models.UserBasic{})
	// db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.GroupBasic{})
	db.AutoMigrate(&models.Contact{})

	//Create
	// user := &models.UserBasic{}
	// user.Name = "Lee"
	// db.Create(user)

	//Read
	// fmt.Println(db.First(user, 1))
	// db.Model(user).Update("PassWord", "4567")

}
