package validation

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
}

func (p *Pagination) LastPage() bool {
	return p.CurrentPage*p.PerPage > p.Total
}
