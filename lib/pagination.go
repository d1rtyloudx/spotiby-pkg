package lib

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"math"
	"strconv"
)

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
	TotalPages  uint64 `json:"total_pages"`
	HasMore     bool   `json:"has_more"`
}

func NewPaginationResponse(totalCount uint64, limit uint64, currentPage uint64) PaginationResponse {
	totalPages := uint64(math.Ceil(float64(totalCount) / float64(limit)))

	return PaginationResponse{
		Limit:       limit,
		CurrentPage: currentPage,
		TotalPages:  totalPages,
		HasMore:     currentPage < totalPages,
	}
}

func ExtractPageQueryParams(ctx echo.Context) (PaginationQuery, error) {
	page, err := strconv.ParseUint(ctx.QueryParam("page"), 10, 64)
	if err != nil {
		return PaginationQuery{}, err
	}

	limit, err := strconv.ParseUint(ctx.QueryParam("limit"), 10, 64)
	if err != nil {
		return PaginationQuery{}, err
	}

	pageQuery := PaginationQuery{
		Page:  page,
		Limit: limit,
	}

	err = pageQuery.Validate()
	if err != nil {
		return PaginationQuery{}, err
	}

	return pageQuery, nil
}
