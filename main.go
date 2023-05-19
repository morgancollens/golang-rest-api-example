package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Driver struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
	Team   string `json:"team"`
}

type Drivers struct {
	Drivers []Driver `json:"drivers"`
}

func _readDriversFromFile() []Driver {
	jsonFile, err := os.Open("data/drivers.json")
	defer jsonFile.Close()

	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var drivers Drivers
	json.Unmarshal(byteValue, &drivers)

	return drivers.Drivers
}

/*
*

	GET /drivers
	Returns a list of drivers participating in the 2023 Formula 1 Season

*
*/
func getDrivers(c *gin.Context) {
	driverList := _readDriversFromFile()

	c.IndentedJSON(http.StatusOK, driverList)
}

/*
*

	GET /drivers/:driverName
	Returns a single driver participating in the 2023 Formula 1 Season,
	when provided with the drivers name.

*
*/
func getSingleDriver(c *gin.Context) {
	driverList := _readDriversFromFile()
	driverName, _ := c.Params.Get("driverName")

	driver := findDriverByName(driverName, driverList)

	c.IndentedJSON(http.StatusOK, driver)
}

func main() {
	router := gin.Default()

	router.GET("/drivers", getDrivers)
	router.GET("/drivers/:driverName", getSingleDriver)

	router.Run("localhost:8080")
}
