package dao

import (
	"{{ .AppName }}/db"
	"{{ .AppName }}/lib/query"
)

func Get(m interface{}) error {
	return db.GetDB().Take(m).Error
}

func Create(m interface{}) error {
	return db.GetDB().Create(m).Error
}

func Update(m interface{}) error {
	return db.GetDB().Save(m).Error
}

func Delete(m interface{}) error {
	return db.GetDB().Delete(m).Error
}

func PageQuery[T any](m *T, p *query.Params) (list []T, totalCount int64, err error) {
	sdb := db.GetDB()

	err = p.QueryCount(sdb).Model(m).Count(&totalCount).Error
	if err != nil {
		return
	}
	if totalCount == 0 {
		return
	}
	err = p.Query(sdb).Find(&list).Error
	return
}
