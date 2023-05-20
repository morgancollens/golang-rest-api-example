package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Driver struct {
	DriverID        string `json:"driverId"`
	PermanentNumber string `json:"permanentNumber"`
	Code            string `json:"code"`
	URL             string `json:"url"`
	GivenName       string `json:"givenName"`
	FamilyName      string `json:"familyName"`
	DateOfBirth     string `json:"dateOfBirth"`
	Nationality     string `json:"nationality"`
}

type DriverTable struct {
	Year    int      `json:"season"`
	Drivers []Driver `json:"Drivers"`
}

type GetAllDriversBody struct {
	DriverTable DriverTable `json:"DriverTable"`
}

type GetAllDriversResponse struct {
	Body GetAllDriversBody `json:"MRData"`
}

var apiEndpoint string = "https://ergast.com/api/f1/2023"

func _getAllDrivers() GetAllDriversResponse {
	resp, err := http.Get(apiEndpoint + "/drivers.json")

	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(resp.Body)

	var data GetAllDriversResponse
	json.Unmarshal(byteValue, &data)

	return data
}

/*
*

	GET /drivers
	Returns a list of drivers participating in the 2023 Formula 1 Season

*
*/
func getDrivers(c *gin.Context) {
	data := _getAllDrivers()

	c.IndentedJSON(http.StatusOK, data.Body.DriverTable.Drivers)
}

/*
*

	GET /drivers/:driverName
	Returns a single driver participating in the 2023 Formula 1 Season,
	when provided with the drivers name.

*
*/
func getSingleDriver(c *gin.Context) {
	driverList := _getAllDrivers()
	driverName, _ := c.Params.Get("driverName")
	driver := findDriverByName(driverName, driverList.Body.DriverTable.Drivers)

	c.IndentedJSON(http.StatusOK, driver)
}

func main() {
	router := gin.Default()

	router.GET("/drivers", getDrivers)
	router.GET("/drivers/:driverName", getSingleDriver)

	router.Run("localhost:8080")
}
