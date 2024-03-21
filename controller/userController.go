package main

import (
	"echoApiRest/models"
	"echoApiRest/services"
	"fmt"
	"github.com/labstack/echo/v4"
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

	// GET Method for Users
	e.GET("/api/users", func(c echo.Context) error {
		// Getting Users
		users, err := services.ReadService()
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, users)
	})

	// POST Method for Users
	e.POST("/api/users", func(c echo.Context) error {
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
	})

	e.PUT("/api/users/:id", func(c echo.Context) error {
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
	})

	e.DELETE("/api/users", func(c echo.Context) error {
		userID := c.QueryParam("id")
		if err := services.DeleteService(userID); err != nil {
			return err
		}
		return c.String(http.StatusOK, "User deleted successfully")
	})

	// Starting Echo Server
	e.Logger.Print(fmt.Sprintf("Listening port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))
}
