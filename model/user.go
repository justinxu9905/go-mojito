package model

import (
	"mojito/database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(user *User) (err error) {
	if err = database.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	fmt.Println(user)
	database.DB.Save(user)
	return nil
}

func GetUser(user *User, id string) (err error) {
	if err := database.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	database.DB.Where("id = ?", id).Delete(user)
	return nil
}