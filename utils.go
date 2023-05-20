package main

import (
	"reflect"
)

/*
*

	When provided with a slice, return a value from that slice based
	on specified criteria

*
*/
func find(slice interface{}, callback func(interface{}) bool) interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("find: not a slice")
	}

	for i := 0; i < s.Len(); i++ {
		value := s.Index(i).Interface()
		if callback(value) {
			return value
		}
	}

	return nil
}

/*
*

	When provided with a name and list of drivers, find a driver by their name
	and return that value

*
*/
func findDriverByName(name string, driverList []Driver) interface{} {
	result := find(driverList, func(p interface{}) bool {
		driver := p.(Driver)
		return (driver.GivenName + " " + driver.FamilyName) == name
	})

	if result != nil {
		return result.(Driver)
	}

	return nil
}
