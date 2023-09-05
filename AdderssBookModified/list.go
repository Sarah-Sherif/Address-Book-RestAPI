package main

import (
	"github.com/gin-gonic/gin"
)

func ListPeople(c *gin.Context) ([]People, error) {
	var peopleDetails []People
	rows, err := DB1.Raw("select * from people").Rows()
	if err != nil {
		print(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var record People
		if err := rows.Scan(&record.PhoneNumber, &record.FirstName, &record.LastName); err != nil {
			print(err.Error())

		}

		peopleDetails = append(peopleDetails, record)

		if err = rows.Err(); err != nil {
			print(err.Error())

		}

	}
	return peopleDetails, err
}

func ListAddress(c *gin.Context) ([]Address, error) {
	var addressDetails []Address
	rows, err := DB1.Raw("select * from address ").Rows()
	if err != nil {
		print(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var record Address
		if err := rows.Scan(&record.Building_id, &record.Pno, &record.StreetName, &record.StreetNo, &record.ZipCode, &record.City, &record.State); err != nil {
			print(err.Error())

		}

		addressDetails = append(addressDetails, record)

		if err = rows.Err(); err != nil {
			print(err.Error())

		}

	}
	return addressDetails, err
}
