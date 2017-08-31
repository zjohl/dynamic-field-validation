package validation

type Pagination struct {
	CurrentPage int
	PerPage     int
	Total       int
}

func (p *Pagination) LastPage() bool {
	return p.CurrentPage*p.PerPage > p.Total
}
