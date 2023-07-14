package util

type paginationInterface interface {
	PageMeta(uint32, uint32) pageMeta
}

type paginationStruct struct{}

func NewPagination() paginationInterface {
	return &paginationStruct{}
}

type pageMeta struct {
	Page   uint32
	Limit  uint32
	Offset uint32
}

func (p *paginationStruct) PageMeta(page, limit uint32) pageMeta {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 1
	}
	offset := limit * (page - 1)
	pageMeta := pageMeta{
		Page:   page,
		Limit:  limit,
		Offset: offset,
	}
	return pageMeta
}
