package services

import (
	"echoApiRest/models"
	"echoApiRest/repositories"
)

func CreateService(user models.User) error {
	err := repositories.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func ReadService() (models.Users, error) {
	users, err := repositories.Read()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func UpdateService(user models.User, userId string) error {
	err := repositories.Update(user, userId)
	if err != nil {
		return err
	}
	return nil
}

func DeleteService(userId string) error {
	err := repositories.Delete(userId)
	if err != nil {
		return err
	}
	return nil
}
