package services

import "hourManager/src/models"

func UserServiceUserGetByName(userName string) (*models.ComUser, error) {
	user, err := models.GetComUserByName(userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserServiceUserUpdate(user *models.ComUser, fields ...string) error {
	err := models.UpdateComUserById(user)
	return err
}

func UserServiceUserGetById(id string) (*models.ComUser, error) {
	user, err := models.GetComUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UserServiceDeleteById(id string) error {

	err := models.DeleteComUser(id)
	if err != nil {
		return err
	}
	return nil

}
