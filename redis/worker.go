package redis

import (
	"encoding/json"
	"mojito/database"
)

type Lost struct {
	ID uint 			`json:"id"`
	Description string 	`json:"description"`
	UserID string 	`json:"user id"`
}

func HandleTask(task string, value string) {
	switch task {
	case "CreateLost":
		b := []byte(value)
		lost := &Lost{}
		err := json.Unmarshal(b, lost)
		if err != nil {
			panic(err)
		}

		if err := database.DB.Create(lost).Error; err != nil {
			panic(err)
		}
	}
}