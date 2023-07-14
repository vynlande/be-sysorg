package request

// program
type CreateExample struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

type ListExample struct {
	Page   uint32 `query:"page"`
	Limit  uint32 `query:"limit"`
	Search string `query:"search"`
	Order  string `query:"order"`
}

func (l ListExample) Init() ListExample {
	l.Page = 1
	l.Limit = 10
	l.Order = "desc"
	return l
}

type PutExample struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

type PatchExample struct {
	Name string `json:"name"`
	Age  int8   `json:"age"`
}
