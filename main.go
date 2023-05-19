package main

import (
	"encoding/json"
	"fmt"
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

/**
*	GET /drivers
*	Returns a list of drivers participating in the 2023 Formula 1 Season
*
**/
func getDrivers(c *gin.Context) {
	fmt.Println("/drivers")
	driverList := _readDriversFromFile()

	c.IndentedJSON(http.StatusOK, driverList)
}

func main() {
	router := gin.Default()

	router.GET("/drivers", getDrivers)

	router.Run("localhost:8080")
}
