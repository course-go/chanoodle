package pagination

type Pagination[T any] struct {
	offset int
	limit  int
}

func New[T any](offset, limit int) Pagination[T] {
	return Pagination[T]{
		offset: offset,
		limit:  limit,
	}
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
