package entity

import "github.com/vynlande/be-sysorg/internal/datastruct"

type CreateExampleReq struct {
	Name string
	Age  int8
}

type ListExampleReq struct {
	Search string
	Limit  uint32
	Page   uint32
	Order  string
}

type ListExampleRes struct {
	Items      []datastruct.Example
	Page       uint32
	Limit      uint32
	TotalRows  uint32
	TotalPages uint32
}

type DetailExampleReq struct {
	ID int
}

type DetailExampleRes struct {
	Item datastruct.Example
}

type DeleteExampleReq struct {
	ID int
}

type PutExampleReq struct {
	ID   int
	Name string
	Age  int8
}

type PatchExampleReq struct {
	ID   int
	Name string
	Age  int8
}
