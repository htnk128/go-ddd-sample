package dto

type PaginationDTO[T comparable] struct {
	Count  int
	Limit  int
	Offset int
	Data   []*T
}

func (p *PaginationDTO[T]) HasMore() bool {
	return p.Count > p.Limit+(p.Limit*p.Offset)
}
