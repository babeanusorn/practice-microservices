package dbhandler

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func Init() *gorm.DB {
	dsn := "host=postgres user=admin password=1234 dbname=practice"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&Person{})

	return db
}
