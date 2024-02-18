package data

import (
	"strings"

	"github.com/arvindeva/touhouapi-cms/internal/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

func (filters Filters) sortColumn() string {
	for _, safeValue := range filters.SortSafeList {
		if filters.Sort == safeValue {
			return strings.TrimPrefix(filters.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + filters.Sort)
}

func (filters Filters) sortDirection() string {
	if strings.HasPrefix(filters.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func (filters Filters) limit() int {
	return filters.PageSize
}

func (filters Filters) offset() int {
	return (filters.Page - 1) * filters.PageSize
}

func ValidateFilters(v *validator.Validator, filters Filters) {
	v.Check(filters.Page > 0, "page", "must be greater than zero")
	v.Check(filters.Page < 10000000, "page", "must be a realistic page number")
	v.Check(filters.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(filters.PageSize <= 100, "page_size", "must not exceed 100")
	v.Check(validator.PermittedValue(filters.Sort, filters.SortSafeList...), "sort", "invalid sort value")
}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

func calculateMetadata(totalRecords, page, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}
	return Metadata{
		CurrentPage:  page,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     (totalRecords + pageSize - 1) / pageSize,
		TotalRecords: totalRecords,
	}
}
