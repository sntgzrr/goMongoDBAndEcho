package main

import (
	"echoApiRest/models"
	"echoApiRest/services"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"os"
	"time"
)

func main() {
	// Setting number port by a Key Environment
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}
	// Creating echo var
	e := echo.New()
	// Using group to create a new router group
	g := e.Group("/api")
	// This middleware logs the server interaction
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[method=${method}, host=${host}, uri=${uri}, status=${status}]\n",
	}))
	// This middleware causes there to be an authentication for use /api router
	g.Use(middleware.BasicAuth(func(username string, password string, context echo.Context) (bool, error) {
		if username == "santiago" && password == "root" {
			return true, nil
		}
		return false, nil
	}))
	// GET Method for Users
	g.GET("/users", getUsers)
	// POST Method for Users
	g.POST("/users", postUser)
	// PUT Method for Users
	g.PUT("/users/:id", putUser)
	// DELETE Method for Users
	g.DELETE("/users", deleteUser)
	// Starting Echo Server
	e.Logger.Print(fmt.Sprintf("Listening port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}

func getUsers(c echo.Context) error {
	// Getting Users
	users, err := services.ReadService()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func postUser(c echo.Context) error {
	var user models.User
	var uid primitive.ObjectID
	// Request Body into User Model
	if err := c.Bind(&user); err != nil {
		return err
	}
	// POST ID and CreatedAt
	user.ID = uid
	user.CreatedAt = time.Now()
	// Creating User
	if err := services.CreateService(user); err != nil {
		return err
	}
	return c.String(http.StatusCreated, "User created successfully")
}

func putUser(c echo.Context) error {
	var user models.User
	userID := c.Param("id")
	// Request Body into User Model
	if err := c.Bind(&user); err != nil {
		return err
	}
	// Updating User
	if err := services.UpdateService(user, userID); err != nil {
		return err
	}
	return c.String(http.StatusOK, "User updated successfully")
}

func deleteUser(c echo.Context) error {
	userID := c.QueryParam("id")
	if err := services.DeleteService(userID); err != nil {
		return err
	}
	return c.String(http.StatusOK, "User deleted successfully")
}
