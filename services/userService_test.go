package services

import (
	"echoApiRest/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var userId string

func TestCreateService(t *testing.T) {
	uid := primitive.ObjectID{}
	userId = uid.Hex()
	user := models.User{
		ID:        uid,
		Name:      "Santiago",
		Email:     "santiago@gmail.com",
		CreatedAt: time.Now(),
	}
	err := CreateService(user)
	if err != nil {
		t.Error("Creating User failed")
	} else {
		t.Log("Creating User passed")
	}
}

func TestReadService(t *testing.T) {
	users, err := ReadService()
	if err != nil {
		t.Error("Reading Users failed")
		t.Fail()
	}
	if len(users) == 0 {
		t.Error("Reading Users don't returned data")
		t.Fail()
	} else {
		t.Log("Reading Users passed")
	}
}

func TestUpdateService(t *testing.T) {
	user := models.User{
		Name:  "Santiago Lozano",
		Email: "santiago777@gmail.com",
	}
	err := UpdateService(user, userId)
	if err != nil {
		t.Error("Updating User failed")
		t.Fail()
	} else {
		t.Log("Updating User passed")
	}
}

func TestDeleteService(t *testing.T) {
	err := DeleteService(userId)
	if err != nil {
		t.Error("Deleting User failed")
		t.Fail()
	} else {
		t.Log("Deleting User passed")
	}
}
