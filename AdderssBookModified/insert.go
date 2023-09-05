package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertPeople(c *gin.Context) {
	var userDetails People
	if err := c.ShouldBind(&userDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"User Details": userDetails,
		})
	}

	query := "Insert into People set phoneNumber=?, firstName=?, lastName=?"
	DB1.Exec(query, userDetails.PhoneNumber, userDetails.FirstName, userDetails.LastName)

}
func InsertAddress(c *gin.Context) {
	var addressDetails Address
	if err := c.ShouldBind(&addressDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"Address Details": addressDetails,
		})
	}

	query := "Insert into Address set BuildingID=?, PNO=?, StreetName=?, StreetNo=?, ZipCode=?, City=?, State=?"
	DB1.Exec(query, addressDetails.Building_id, addressDetails.Pno, addressDetails.StreetName, addressDetails.StreetNo, addressDetails.ZipCode, addressDetails.City, addressDetails.State)

}
