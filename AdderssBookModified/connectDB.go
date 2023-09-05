package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type People struct {
	PhoneNumber string `json:"phoneNo" form:"phoneNo" gorm:"primary_key" binding:"required"` //binding when entering the data is a must
	FirstName   string `json:"fName" form:"fName" binding:"required"`
	LastName    string `json:"lName" form:"lName"`
}
type Address struct {
	Building_id string `json:"building_id" form:"building_id" gorm:"primary_key" binding:"required"`
	Pno         string `json:"pno" form:"pno" binding:"required"`
	StreetName  string `json:"streetName" form:"streetName"`
	StreetNo    string `json:"streetNo" form:"streetNo" `
	ZipCode     string `json:"zipCode" form:"zipCode"`
	City        string `json:"city" form:"city"`
	State       string `json:"state" form:"state"`
}

var DB1 *gorm.DB

func connectDB() {
	dsn := "Sarah:S04770371~@tcp(localhost:3306)/addressbook"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to the DB!")
	}

	DB1 = database //will be used in the controller to get access to the database
	print("hello")
}
