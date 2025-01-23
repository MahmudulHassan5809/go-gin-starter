package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)


type QueryService struct{}

type PaginationOptions struct {
	PageSize int `json:"page_size"`
	Page     int `json:"page"`
	Skip     int `json:"skip"`
}

type FilterOptions struct {
	Search     string
	SearchFields []string                 
	Filters    map[string]interface{} 
	Sorting    map[string]string      
	Pagination *PaginationOptions 
}


const DefaultPageSize = 10
const DefaultPage = 1


func (filterOptions *FilterOptions) SetDefaults() {
	if filterOptions.Pagination.PageSize == 0 {
		filterOptions.Pagination.PageSize = DefaultPageSize
	}
	if filterOptions.Pagination.Page == 0 {
		filterOptions.Pagination.Page = DefaultPage
	}
	if filterOptions.Pagination.Skip == 0 {
		filterOptions.Pagination.Skip = (filterOptions.Pagination.Page - 1) * filterOptions.Pagination.PageSize
	}
}

func NewQueryService() *QueryService {
	return &QueryService{}
}


func (q *QueryService) ParseQueryParams(c *gin.Context, filterFields []string, searchFields []string) (*FilterOptions, error) {
	search := c.DefaultQuery("search", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", fmt.Sprintf("%d", DefaultPage)))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", fmt.Sprintf("%d", DefaultPageSize)))
	sortField := c.DefaultQuery("sort_field", "")
	sortOrder := c.DefaultQuery("sort_order", "asc")
	filterParams := c.DefaultQuery("filters", "")
	filters := map[string]interface{}{}
	if filterParams != "" {
		err := json.Unmarshal([]byte(filterParams), &filters)
		if err != nil {
			return nil, errors.New("invalid filters format")
		}
		allowedFilters := map[string]interface{}{}
		for _, field := range filterFields {
			if value, exists := filters[field]; exists {
				allowedFilters[field] = value
			}
		}
		filters = allowedFilters
	}
	skip := (page - 1) * pageSize
	filterOptions := &FilterOptions{
		Search:  search,
		SearchFields: searchFields,
		Filters: filters,
		Sorting: map[string]string{
			sortField: sortOrder,
		},
		Pagination: &PaginationOptions{
			Page:     page,
			PageSize: pageSize,
			Skip:     skip,
		},
	}
	return filterOptions, nil
}
