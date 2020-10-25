package model

import (
	"mojito/database"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateFound(found *Found) (err error) {
	if err = database.DB.Create(found).Error; err != nil {
		return err
	}
	return nil
}

func UpdateFound(found *Found, id string) (err error) {
	fmt.Println(found)
	database.DB.Save(found)
	return nil
}

func GetFound(found *Found, id string) (err error) {
	if err := database.DB.Where("id = ?", id).First(found).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFound(found *Found, id string) (err error) {
	database.DB.Where("id = ?", id).Delete(found)
	return nil
}