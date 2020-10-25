package model

import (
	"encoding/json"
	"mojito/database"
	"mojito/redis"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func CreateLost(lost *Lost) (err error) {
	bi, err := json.Marshal(&lost)
	if err != nil {
		panic(err)
	}
	err = redis.RB.SendTask("CreateLost", string(bi))
	/*if err = database.DB.Create(lost).Error; err != nil {
		return err
	}
	return nil*/
	return err
}

func UpdateLost(lost *Lost, id string) (err error) {
	fmt.Println(lost)
	database.DB.Save(lost)
	return nil
}

func GetLost(lost *Lost, id string) (err error) {
	if err := database.DB.Where("id = ?", id).First(lost).Error; err != nil {
		return err
	}
	return nil
}

func DeleteLost(lost *Lost, id string) (err error) {
	database.DB.Where("id = ?", id).Delete(lost)
	return nil
}