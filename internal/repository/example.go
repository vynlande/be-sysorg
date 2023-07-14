package repository

import (
	"database/sql"
	"fmt"

	"github.com/vynlande/be-sysorg/internal/datastruct"
	"github.com/vynlande/be-sysorg/internal/serrepconnector"
	"github.com/vynlande/be-sysorg/pkg/util"
	"gorm.io/gorm"
)

type ExampleQuery interface {
	ListExample(serrepconnector.ListExampleReq) *sql.Rows
	CreateExample(serrepconnector.CreateExampleReq)
	DetailExample(serrepconnector.DetailExampleReq) *sql.Rows
	PutExample(serrepconnector.PutExampleReq)
	PatchExample(serrepconnector.PatchExampleReq)
	DeleteExample(serrepconnector.DeleteExampleReq)
}

type exampleQuery struct {
	gormDB *gorm.DB
	sqlDB  *sql.DB
}

func (e *exampleQuery) ListExample(req serrepconnector.ListExampleReq) *sql.Rows {
	query := fmt.Sprintf(`select id, uuid, name, age, created_at, updated_at, count(*) over() as total_rows from example 
		where lower(name) like lower('%%%v%%') order by id %v limit %v offset %v`,
		req.Search, req.Order, req.Limit, req.Offset)
	rows, err := e.sqlDB.Query(query)
	util.NewError().Panic(err)
	return rows
}

func (e *exampleQuery) CreateExample(req serrepconnector.CreateExampleReq) {
	err := e.gormDB.Create(&req.Item).Error
	util.NewError().Panic(err)
}

func (e *exampleQuery) DetailExample(req serrepconnector.DetailExampleReq) *sql.Rows {
	query := fmt.Sprintf(`
		select id, uuid, name, age, created_at, updated_at from example where id = %v limit 1
	`, req.ID)
	rows, err := e.sqlDB.Query(query)
	util.NewError().Panic(err)
	return rows
}

func (e *exampleQuery) PutExample(req serrepconnector.PutExampleReq) {
	err := e.gormDB.Model(&datastruct.Example{}).Select("*").Where("id = ?", req.Item.ID).Updates(req.Item).Error
	util.NewError().Panic(err)
}

func (e *exampleQuery) PatchExample(req serrepconnector.PatchExampleReq) {
	err := e.gormDB.Model(&datastruct.Example{}).Where("id = ?", req.Item.ID).Updates(req.Item).Error
	util.NewError().Panic(err)
}

func (e *exampleQuery) DeleteExample(req serrepconnector.DeleteExampleReq) {
	err := e.gormDB.Delete(&datastruct.Example{}, req.ID).Error
	util.NewError().Panic(err)
}
