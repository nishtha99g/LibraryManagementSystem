// models/database.go
package models

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitDatabase() {
    dsn := "root:Nishtha@paytm21@tcp(127.0.0.1:3306)/books?parseTime=true"
    connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to the database")
    }

    DB = connection
    AutoMigrateModels()
}

func AutoMigrateModels() {
    DB.AutoMigrate(&Book{})
    DB.AutoMigrate(&User{})
    DB.AutoMigrate(&Claims{})
}
