package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PeopleReturn(c *gin.Context) {
	recordPeople, err := ListPeople(c)
	if err != nil {
		print("Error %s when selecting record ", err)
	}
	for _, record := range recordPeople {
		c.JSON(http.StatusOK, gin.H{"People Details": record})

	}

}
func AddressReturn(c *gin.Context) {
	recordAddress, err := ListAddress(c)
	if err != nil {
		print("Error %s when selecting record ", err)
	}
	for _, record := range recordAddress {
		c.JSON(http.StatusOK, gin.H{"Address Details": record})

	}

}
