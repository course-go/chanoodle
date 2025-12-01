package pagination

import "errors"

var ErrInvalidPagination = errors.New("invalid pagination parameters")

const defaultLimit = 250

// Pagination represents pagination parameters.
type Pagination[T any] struct {
	limit  int
	offset int
}

func New[T any](limit, offset int) (p Pagination[T], err error) {
	if limit < 0 || offset < 0 {
		return Pagination[T]{}, ErrInvalidPagination
	}

	if limit == 0 {
		limit = defaultLimit
	}

	return Pagination[T]{
		limit:  limit,
		offset: offset,
	}, nil
}

func (p *Pagination[T]) Paginate(slice []T) []T {
	if p.limit == 0 || p.offset >= len(slice) {
		return make([]T, 0)
	}

	from := p.offset
	to := min(from+p.limit, len(slice))

	return slice[from:to]
}
