package database

import (
    "app/config"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var d *gorm.DB

// Init initializes database
func Init(isReset bool, models ...interface{}) {
    c := config.GetConfig()
    var err error
    dns := c.GetString("db.url")
    d, err = gorm.Open(mysql.Open(dns))

    if err != nil {
        panic(err)
    }

    d.AutoMigrate(models...)
}

// GetDB returns database connection
func GetDB() *gorm.DB {
    return d
}
