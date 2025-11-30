package pagination

import "errors"

var ErrInvalidPagination = errors.New("invalid pagination parameters")

type Pagination[T any] struct {
	limit  int
	offset int
}

func New[T any](limit, offset int) (p Pagination[T], err error) {
	if limit < 0 || offset < 0 {
		return Pagination[T]{}, ErrInvalidPagination
	}

	return Pagination[T]{
		limit:  limit,
		offset: offset,
	}, nil
}

func (p *Pagination[T]) Paginate(slice []T) []T {
	if p.offset >= len(slice) || p.limit == 0 {
		return make([]T, 0)
	}

	if p.offset > 0 && p.offset < len(slice) {
		return slice[p.offset:]
	}

	if p.limit > 0 && p.limit < len(slice) {
		return slice[:p.limit]
	}

	return slice
}
