package models
import (
	"time"
)

type Common struct {
	ID		uint64 `json:"id,omitempty"  gorm:"unique;not null" form:"id"`
	CreatedAt    time.Time `json:",omitempty"`
	UpdatedAt    time.Time `json:",omitempty"`
}

var models = [] interface{}{ 
	&Restaurant{},
}

// GetConfig returns config
func GetModels() []interface{} {
    return models
}