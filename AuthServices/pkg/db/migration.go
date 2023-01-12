package db

import (
	"fmt"
	"github.com/man-droid23/simple-microservices/pkg/models/entity"
)

func Migration() {
	err := DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic("Failed to migrate database")
	}
	fmt.Println("Database migrated successfully")
}
