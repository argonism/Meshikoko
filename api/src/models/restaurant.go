package models
import (
	"app/database"
)

type Restaurant struct {
	Common
	Name    string		`json:"name" gorm:"unique;not null" form:"name"`
}

func (u *Restaurant) Create() (err error) {
    db := database.GetDB()
    return db.Create(u).Error
}

func (u *Restaurant) FindByID(id uint64) (err error) {
    db := database.GetDB()
    err = db.First(u, id).Error
    return
}
