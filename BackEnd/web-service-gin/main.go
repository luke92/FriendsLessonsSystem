package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Domain
type user struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

//Data
var users = []user{
	{ID: "1", Username: "luke92", Password: "mypassword", Email: "lucasjv92@gmail.com", Name: "Lucas", Lastname: "Vargas"},
	{ID: "2", Username: "henry", Password: "henryPass", Email: "henry@modak.live", Name: "Henry", Lastname: "Canastero"},
}

//Handlers
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of users, looking for
	// an user whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func getFriendships(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getFriendsByUserId(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of users, looking for
	// an user whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func getLessonsByUserId(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of users, looking for
	// an user whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserById)
	router.GET("/friendships", getFriendships)
	router.GET("/users/:id/friends", getFriendsByUserId)
	router.GET("/users/:id/lessons", getLessonsByUserId)

	router.Run("localhost:8080")
}
