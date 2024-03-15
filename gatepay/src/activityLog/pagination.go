package activityLog

import "gorm.io/gorm"

type Pagination struct {
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
	SortBy   string  `json:"sort_by"`
	Filters  Filters `json:"filters"`
}

func (p *Pagination) Scopes() []func(*gorm.DB) *gorm.DB {
	return []func(*gorm.DB) *gorm.DB{
		p.paginationScope,
		p.filtersScope,
	}
}

func (p *Pagination) paginationScope(db *gorm.DB) *gorm.DB {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.PageSize == 0 {
		p.PageSize = 10
	}

	return db.Offset((p.Page - 1) * p.PageSize).Limit(p.PageSize)
}

func (p *Pagination) filtersScope(db *gorm.DB) *gorm.DB {
	if len(p.Filters) == 0 {
		return db
	}

	for i, f := range p.Filters {
		if i == 0 || f.LogicalOperator == "and" {
			db = db.Where(f.query(), f.Value)
			continue
		}

		db = db.Or(f.query(), f.Value)
	}

	return db
}
