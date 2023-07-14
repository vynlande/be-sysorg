package datastruct

import "time"

type Example struct {
	ID        int       `gorm:"column:id" json:"-"`
	UUID      string    `gorm:"-" db:"uuid" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Age       int8      `gorm:"column:age" json:"age"`
	CreatedAt time.Time `gorm:"column:created_at; default:current_timestamp" db:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at; default:current_timestamp" db:"updated_at" json:"updated_at"`
	TotalRows uint32    `gorm:"-" db:"total_rows" json:"-"`
}

func (Example) TableName() string {
	return "example"
}
