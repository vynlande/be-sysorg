package service

import (
	"math"

	"github.com/blockloop/scan/v2"
	"github.com/vynlande/be-sysorg/internal/datastruct"
	"github.com/vynlande/be-sysorg/internal/entity"
	"github.com/vynlande/be-sysorg/internal/repository"
	"github.com/vynlande/be-sysorg/internal/serrepconnector"
	"github.com/vynlande/be-sysorg/pkg/util"
)

type ExampleService interface {
	CreateExample(entity.CreateExampleReq)
	ListExample(entity.ListExampleReq) entity.ListExampleRes
	DetailExample(entity.DetailExampleReq) entity.DetailExampleRes
	PutExample(entity.PutExampleReq)
	PatchExample(entity.PatchExampleReq)
	DeleteExample(entity.DeleteExampleReq)
}

type exampleService struct {
	dao repository.DAO
}

func NewExampleService(dao repository.DAO) ExampleService {
	return &exampleService{
		dao: dao,
	}
}

func (e *exampleService) CreateExample(req entity.CreateExampleReq) {
	item := datastruct.Example{
		Name: req.Name,
		Age:  req.Age,
	}
	e.dao.NewExampleQuery().CreateExample(serrepconnector.CreateExampleReq{
		Item: item,
	})
}

func (e *exampleService) ListExample(req entity.ListExampleReq) (res entity.ListExampleRes) {
	pageMeta := util.NewPagination().PageMeta(req.Page, req.Limit)
	sqlRows := e.dao.NewExampleQuery().ListExample(serrepconnector.ListExampleReq{
		Search: req.Search,
		Order:  req.Order,
		Limit:  pageMeta.Limit,
		Offset: pageMeta.Offset,
	})
	err := scan.Rows(&res.Items, sqlRows)
	util.NewError().Panic(err)
	res.Page = pageMeta.Page
	res.Limit = pageMeta.Limit
	if len(res.Items) > 0 {
		res.TotalRows = res.Items[0].TotalRows
	}
	res.TotalPages = uint32(math.Ceil(float64(res.TotalRows) / float64(pageMeta.Limit)))
	return
}

func (e *exampleService) DetailExample(req entity.DetailExampleReq) (res entity.DetailExampleRes) {
	sqlRows := e.dao.NewExampleQuery().DetailExample(serrepconnector.DetailExampleReq{
		ID: req.ID,
	})
	err := scan.Row(&res.Item, sqlRows)
	util.NewError().Panic(err)
	return
}

func (e *exampleService) PutExample(req entity.PutExampleReq) {
	item := datastruct.Example{
		ID:   req.ID,
		Name: req.Name,
		Age:  req.Age,
	}
	e.dao.NewExampleQuery().PutExample(serrepconnector.PutExampleReq{
		Item: item,
	})
}

func (e *exampleService) PatchExample(req entity.PatchExampleReq) {
	item := datastruct.Example{
		ID:   req.ID,
		Name: req.Name,
		Age:  req.Age,
	}
	e.dao.NewExampleQuery().PatchExample(serrepconnector.PatchExampleReq{
		Item: item,
	})
}

func (e *exampleService) DeleteExample(req entity.DeleteExampleReq) {
	e.dao.NewExampleQuery().DeleteExample(serrepconnector.DeleteExampleReq{
		ID: req.ID,
	})
}
