package serrepconnector

import "github.com/vynlande/be-sysorg/internal/datastruct"

type CreateExampleReq struct {
	Item datastruct.Example
}

type ListExampleReq struct {
	Search string
	Limit  uint32
	Offset uint32
	Order  string
}

type DetailExampleReq struct {
	ID int
}

type DeleteExampleReq struct {
	ID int
}

type PutExampleReq struct {
	Item datastruct.Example
}

type PatchExampleReq struct {
	Item datastruct.Example
}
