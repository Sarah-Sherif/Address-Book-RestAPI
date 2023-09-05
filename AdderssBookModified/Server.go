package main

import (
	"github.com/gin-gonic/gin"
)

func ServerCreation() {
	router := gin.Default() //router is the server
	router.SetTrustedProxies(nil)
	router.POST("/userDetails", InsertPeople) //endpoint is userDetails which is handler created to recieve requests(path that will be appended to localhos:50 URL)
	router.POST("/addressDetails", InsertAddress)
	router.GET("/peopleDetails", PeopleReturn)
	router.GET("/addressDetails", AddressReturn)
	router.GET("/personObj", FilterRecordJsonReturn)
	router.Run("localhost:50") //where localhost is the path and 50 is the port

}
