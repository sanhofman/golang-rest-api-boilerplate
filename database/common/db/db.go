package db

import (
    "log"

	"github.com/wpcodevo/golang-mongodb/database/common/dbModels"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func Init(url string) *gorm.DB {
    db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&dbModels.User{})
    db.AutoMigrate(&dbModels.Child{})

    return db
}
