// I LOOK HERE https://go.dev/doc/tutorial/web-service-gin
// AND I DID NOT TOTALY UNDERSTAND BUT OKAY

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// user represents data about a record user.
type user struct {
	ID    string `json:"id"`
	Title string `json:"user"`
}

// users slice to seed record album data.
var users = []user{
	{ID: "1", Title: "Ali"},
	{ID: "2", Title: "Veli"},
	{ID: "3", Title: "Melih"},
}

func main() {
	router := gin.Default()
	router.GET("/user", getUser)
	router.GET("/user/:id", getUserByID)
	router.POST("/user", postUser)
	router.DELETE("/user/:id", deleteUserByID)
	router.PUT("/user/:id", updateUserByID)
	router.Run("localhost:8081")
}

// getUser responds with the list of all albums as JSON.
func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

// postUser adds an album from JSON received in the request body.
func postUser(c *gin.Context) {
	var newUser user

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new album to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// getUserByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func deleteUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {

			arrayId, _ := strconv.ParseInt(a.ID, 10, 64)

			users = append(users[:arrayId-1], users[arrayId:]...)
			fmt.Println()
			fmt.Println(users)
			fmt.Println()

			c.IndentedJSON(http.StatusOK, `{data:"done"}`)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateUserByID(c *gin.Context) {
	id := c.Param("id")
	var newUser user
	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {

			if err := c.BindJSON(&newUser); err != nil {
				return
			}
			arrayId, _ := strconv.ParseInt(a.ID, 10, 64)

			users[arrayId-1] = newUser
			fmt.Println(users[arrayId-1])
			c.IndentedJSON(http.StatusOK, `{data:"done"}`)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
