package model

type Paginate struct {
	Total     int64
	Page      int
	NumOfPage int64
	Content   interface{}
}
