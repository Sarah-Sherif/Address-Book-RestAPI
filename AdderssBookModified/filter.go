package main

import (
	"github.com/gin-gonic/gin"
)

type PersonAddr struct {
	Phone     string `json:"phoneNo" form:"phoneNo" gorm:"primary_key"`
	FName     string `json:"fName" form:"fName" binding:"required"`
	LName     string `json:"lName" form:"lName"`
	B_id      string `json:"b_id" form:"b_id"`
	SAddrName string `json:"streetName" form:"streetName"`
	SAddrNo   string `json:"streetNo" form:"streetNo"`
	Code      string `json:"zipCode" form:"zipCode"`
	AddrCity  string `json:"city" form:"city"`
	AddrState string `json:"state" form:"state"`
}

func FilterPerson(c *gin.Context, firstName string) ([]PersonAddr, error) {

	var records []PersonAddr
	rows, err := DB1.Table("People").Select("People.PhoneNumber, People.FirstName, People.LastName, Address.BuildingID, Address.StreetName, Address.StreetNo, Address.ZipCode, Address.City, Address.State").Joins("inner join Address on Address.PNO = People.PhoneNumber").Where("People.FirstName = ?", firstName).Rows()
	if err != nil {
		print(err.Error())
	}
	for rows.Next() {
		var recordDetails PersonAddr

		if err := rows.Scan(&recordDetails.Phone, &recordDetails.FName, &recordDetails.LName, &recordDetails.B_id, &recordDetails.SAddrName, &recordDetails.SAddrNo, &recordDetails.Code, &recordDetails.AddrCity, &recordDetails.AddrState); err != nil {
			print(err.Error())
		}

		records = append(records, recordDetails)

	}
	if err = rows.Err(); err != nil {
		return records, err
	}
	return records, nil
}
