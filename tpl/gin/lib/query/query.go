package query

import (
	"gorm.io/gorm"
)

type Params struct {
	Queries []QueryPair // 条件
	OrderBy []OrderBy   // 排序
	Offset  int64
	Limit   int64
}

type QueryPair struct {
	Query string        // 查询
	Args  []interface{} // 参数
	IsOr  bool          //或查询
}

type OrderBy struct {
	Column string
	Asc    bool
}

func (p *Params) Where(query string, args ...interface{}) *Params {
	p.Queries = append(p.Queries, QueryPair{query, args, false})
	return p
}

func (p *Params) OrWhere(query string, args ...interface{}) *Params {
	p.Queries = append(p.Queries, QueryPair{query, args, true})
	return p
}

func (p *Params) Eq(column string, args ...interface{}) *Params {
	p.Where(column+" = ?", args)
	return p
}

func (p *Params) NotEq(column string, args ...interface{}) *Params {
	p.Where(column+" <> ?", args)
	return p
}

func (p *Params) Gte(column string, args ...interface{}) *Params {
	p.Where(column+" >= ?", args)
	return p
}

func (p *Params) Gt(column string, args ...interface{}) *Params {
	p.Where(column+" > ?", args)
	return p
}

func (p *Params) Lte(column string, args ...interface{}) *Params {
	p.Where(column+" <= ?", args)
	return p
}

func (p *Params) Lt(column string, args ...interface{}) *Params {
	p.Where(column+" < ?", args)
	return p
}

func (p *Params) In(column string, args []interface{}) *Params {
	p.Where(column+" IN (?)", args)
	return p
}

func (p *Params) Like(column string, str string) *Params {
	p.Where(column+" like ?", "%"+str+"%")
	return p
}

func (p *Params) OrLike(column string, str string) *Params {
	p.OrWhere(column+" like ?", "%"+str+"%")
	return p
}

func (q *Params) Query(db *gorm.DB) *gorm.DB {
	if len(q.Queries) > 0 {
		for _, pair := range q.Queries {
			if pair.IsOr {
				db = db.Or(pair.Query, pair.Args...)
			} else {
				db = db.Where(pair.Query, pair.Args...)
			}

		}
	}

	if len(q.OrderBy) > 0 {
		for _, item := range q.OrderBy {
			if item.Asc {
				db = db.Order(item.Column + " asc")
			} else {
				db = db.Order(item.Column + " desc")
			}
		}
	}

	if q.Offset > 0 {
		db = db.Offset(int(q.Offset))
	}

	if q.Limit > 0 {
		db = db.Limit(int(q.Limit))
	}

	return db
}

func (q *Params) QueryCount(db *gorm.DB) *gorm.DB {
	if len(q.Queries) > 0 {
		for _, pair := range q.Queries {
			db = db.Where(pair.Query, pair.Args...)
		}
	}

	return db
}
