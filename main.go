package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Person struct defines a person with an ID, Name, Surname, Age, and multiple Addresses.
type Person struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Surname  string   `json:"surname"`
	Age      int      `json:"age"`
	Addresses []string `json:"addresses"` // List of address IDs
}

// Address struct defines an address with an ID, Location, and multiple Persons.
type Address struct {
	ID       string   `json:"id"`
	Location string   `json:"location"`
	Persons  []string `json:"persons"` // List of person IDs
}

// Example data
var persons = []Person{
	{ID: "1", Name: "Zehra", Surname: "Mangal", Age: 23, Addresses: []string{"1", "2"}},
	{ID: "2", Name: "Azra", Surname: "Mangal", Age: 18, Addresses: []string{"1"}},
	{ID: "3", Name: "Zümra", Surname: "Mangal", Age: 2, Addresses: []string{"2"}},
}

var addresses = []Address{
	{ID: "1", Location: "Üsküdar", Persons: []string{"1", "2"}},
	{ID: "2", Location: "Kadiköy", Persons: []string{"1", "3"}},
}

// Get all persons
func getPersons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)
}

// Get all addresses
func getAddresses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, addresses)
}

// Get person by ID
func getPersonByID(c *gin.Context) {
	id := c.Param("id")
	for _, person := range persons {
		if person.ID == id {
			c.IndentedJSON(http.StatusOK, person)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "person not found"})
}

// Get address by ID
func getAddressByID(c *gin.Context) {
	id := c.Param("id")
	for _, address := range addresses {
		if address.ID == id {
			c.IndentedJSON(http.StatusOK, address)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "address not found"})
}

func main() {
	router := gin.Default()
	router.GET("/persons", getPersons)
	router.GET("/persons/:id", getPersonByID)
	router.GET("/addresses", getAddresses)
	router.GET("/addresses/:id", getAddressByID)

	router.Run("localhost:8080")
}

