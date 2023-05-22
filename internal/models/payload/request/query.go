package request

type (
	PaginationRequest struct {
		Page  int64  `json:"page" query:"page"`
		Limit int64  `json:"per_page" query:"per_page"`
		Key   string `json:"key" query:"q"`
	}
)

func (p *PaginationRequest) ValidatePagination() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 {
		p.Limit = 10
	}
}
