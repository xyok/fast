package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type DeletedAt = gorm.DeletedAt
type BaseModel struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt DeletedAt  `gorm:"index" json:"-"`
}

type IntArray []int
type StringArray []string

// Value
func (j IntArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan
func (j *IntArray) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), j)
}

// Value
func (j StringArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan
func (j *StringArray) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), j)
}
