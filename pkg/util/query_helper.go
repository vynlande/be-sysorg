package util

import (
	"fmt"

	"github.com/blockloop/scan/v2"
	"github.com/vynlande/be-sysorg/infrastructure"
)

type queryInterface interface {
	GetIDbyUUID(string, string) int
}

type queryStruct struct{}

func NewQueryHelper() queryInterface {
	return &queryStruct{}
}

func (q *queryStruct) GetIDbyUUID(table string, uuid string) (id int) {
	query := fmt.Sprintf(`
		select id from %v t where uuid = '%v'
	`, table, uuid)
	sqlRows, err := infrastructure.SqlDB.Query(query)
	NewError().Panic(err)
	err = scan.Row(&id, sqlRows)
	NewError().Panic(err)
	return
}
