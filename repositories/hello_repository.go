package repositories

import (
	"fmt"
	"main/config"
	"main/database"
)

func GetAll() (database.HelloTable, error) {
	var response database.HelloTable
	err := config.DB.Take(&response).Error
	if err != nil {
		fmt.Println(err)
		return database.HelloTable{}, err
	}
	return response, nil
}
