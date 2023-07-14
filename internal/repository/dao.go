package repository

import (
	"database/sql"

	"gorm.io/gorm"
)

type DAO interface {
	NewExampleQuery() ExampleQuery
}

type dao struct{}

var GormDB *gorm.DB
var SqlDB *sql.DB

func NewDAO(sqlDB *sql.DB, gormDB *gorm.DB) DAO {
	GormDB = gormDB
	SqlDB = sqlDB
	return &dao{}
}

func (d *dao) NewExampleQuery() ExampleQuery {
	return &exampleQuery{
		gormDB: GormDB,
		sqlDB:  SqlDB,
	}
}
