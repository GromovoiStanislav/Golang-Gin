package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// User структура представляет собой простую модель пользователя
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{}

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Просто для примера, присваиваем временный ID
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	c.JSON(http.StatusCreated, newUser)
}

func getUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Ищем пользователя по ID
	var foundUser *User
	for i := range users {
		if users[i].ID == userID {
			foundUser = &users[i]
			break
		}
	}

	// Если пользователь не найден, возвращаем 404
	if foundUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Возвращаем найденного пользователя
	c.JSON(http.StatusOK, foundUser)
}

func updateUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Ищем пользователя по ID
	var updatedUser *User
	for i := range users {
		if users[i].ID == userID {
			// Найден пользователь, модифицируем его
			if err := c.ShouldBindJSON(&users[i]); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			updatedUser = &users[i]
			break
		}
	}

	// Если пользователь не найден, возвращаем 404
	if updatedUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Возвращаем обновленного пользователя
	c.JSON(http.StatusOK, updatedUser)
}

func deleteUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Ищем индекс пользователя по ID
	userIndex := -1
	for i := range users {
		if users[i].ID == userID {
			userIndex = i
			break
		}
	}

	// Если пользователь не найден, возвращаем 404
	if userIndex == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Удаляем пользователя из списка
	users = append(users[:userIndex], users[userIndex+1:]...)

	// Возвращаем сообщение об успешном удалении
	c.JSON(http.StatusOK, gin.H{"message": "User deleted", "userID": userID})
}


func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	//r := gin.New()
	r := gin.Default()
	

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})

	users := r.Group("/users")
	{
		users.GET("/", listUsers)
		users.POST("/", createUser)
		users.GET("/:id", getUser)
		users.PUT("/:id", updateUser)
		users.DELETE("/:id", deleteUser)
	}

	return r
}

func main() {
	r := setupRouter()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Run(":8080")
}