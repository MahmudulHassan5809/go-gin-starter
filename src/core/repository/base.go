package repository

import (
	"errors"
	"fmt"
	"gin_starter/src/core/services"
	"strings"

	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) GetByField(field, value string) (*T, error) {
	var entity T
	err := r.db.Where(field+" = ?", value).First(&entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &entity, nil
}

func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *BaseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *BaseRepository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}

func (r *BaseRepository[T]) buildFilters(query *gorm.DB, filters map[string]interface{}) *gorm.DB {
	for field, value := range filters {
		fieldParts := strings.Split(field, "__")
		fieldName := fieldParts[0]
		operator := "exact"
		if len(fieldParts) > 1 {
			operator = fieldParts[1]
		}
		switch operator {
		case "exact":
			query = query.Where(fmt.Sprintf("%s = ?", fieldName), value)
		case "in":
			query = query.Where(fmt.Sprintf("%s IN (?)", fieldName), value)
		default:
			query = query.Where(fmt.Sprintf("%s = ?", fieldName), value)
		}
	}
	return query
}

func (r *BaseRepository[T]) buildSorting(query *gorm.DB, sorting map[string]string) *gorm.DB {
	fmt.Println(sorting)
	for field, direction := range sorting {
		if field == "" || (direction != "asc" && direction != "desc") {
			continue
		}
		orderClause := fmt.Sprintf("%s %s", field, strings.ToUpper(direction))
		query = query.Order(orderClause)
	}
	return query
}

func (r *BaseRepository[T]) paginate(query *gorm.DB, pagination *services.PaginationOptions) *gorm.DB {
	if pagination != nil {
		query = query.Offset(pagination.Skip).Limit(pagination.PageSize)
	}
	return query
}

func (r *BaseRepository[T]) buildSearch(query *gorm.DB, search string, searchFields []string) *gorm.DB {
	if search != "" && len(searchFields) > 0 {
		conditions := []string{}
		args := []interface{}{}
		for _, field := range searchFields {
			conditions = append(conditions, fmt.Sprintf("%s LIKE ?", field))
			args = append(args, "%"+search+"%")
		}
		query = query.Where(strings.Join(conditions, " OR "), args...)
	}
	return query
}

func (r *BaseRepository[T]) PaginateAndFilter(filterOptions *services.FilterOptions, entity *T) ([]T, int64, error) {
	filterOptions.SetDefaults()
	query := r.db.Model(entity)
	query = r.buildSearch(query, filterOptions.Search, filterOptions.SearchFields)
	query = r.buildFilters(query, filterOptions.Filters)
	query = r.buildSorting(query, filterOptions.Sorting)
	query = r.paginate(query, filterOptions.Pagination)
	var totalCount int64
	err := query.Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}
	var entities []T
	err = query.Find(&entities).Error
	if err != nil {
		return nil, 0, err
	}
	return entities, totalCount, nil
}
