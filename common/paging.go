package common

type Paging struct {
	Page   int   `json:"page" form:"page"`
	Limit  int   `json:"limit" form:"limit"`
	Total  int64 `json:"total" form:"-"`
	Cursor int   `json:"cursor" form:"cursor"`
}

func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 50
	}
}
