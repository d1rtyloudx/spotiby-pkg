package lib

import "fmt"

type PaginationQuery struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}

func (p *PaginationQuery) Validate() error {
	if p.Page == 0 || p.Limit == 0 {
		return fmt.Errorf("invalid pagination request")
	}

	return nil
}

func (p *PaginationQuery) GetOffset() uint64 {
	if p.Page == 0 {
		return 0
	}

	return (p.Page - 1) * p.Limit
}

type PaginationResponse struct {
	Limit       uint64 `json:"limit"`
	CurrentPage uint64 `json:"current_page"`
	TotalPages  uint64 `json:"total"`
	HasMore     bool   `json:"has_more"`
}

func NewPaginationResponse(totalCount uint64, limit uint64, currentPage uint64) PaginationResponse {
	var totalPages uint64

	total := totalCount / limit

	remainder := total % limit
	if remainder == 0 {
		totalPages = total
	} else {
		totalPages = totalPages + 1
	}

	return PaginationResponse{
		Limit:       limit,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		HasMore:     currentPage < totalPages,
	}
}
