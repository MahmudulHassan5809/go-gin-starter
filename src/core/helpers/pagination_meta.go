package helpers

import "math"


type PaginationMeta struct {
	TotalCount  int64 `json:"total_count"`
	TotalPages  int   `json:"total_pages"`
	CurrentPage int   `json:"current_page"`
	PrevPage    int   `json:"prev_page"`
	NextPage    int   `json:"next_page"`
}


func GeneratePaginationMeta(totalCount int64, currentPage, pageSize int) PaginationMeta {
	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	prevPage := currentPage - 1
	if prevPage < 1 {
		prevPage = 1
	}
	nextPage := currentPage + 1
	if nextPage > totalPages {
		nextPage = totalPages
	}

	return PaginationMeta{
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: currentPage,
		PrevPage:    prevPage,
		NextPage:    nextPage,
	}
}
