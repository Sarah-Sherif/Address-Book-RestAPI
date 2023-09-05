package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FilterRecordJsonReturn(c *gin.Context) { //this function returns the person's address details by his/her name
	name := c.Request.URL.Query().Get("FName")
	if err := c.ShouldBind(&name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	pdetailsRecord, err := FilterPerson(c, name)
	if err != nil {
		print(err.Error())
	}

	for _, personRecordDetail := range pdetailsRecord {
		c.JSON(http.StatusOK, gin.H{"Person's Address Details": personRecordDetail})

	}
}
