package handler

import (
	"net/http"
	"strconv"
	"strings"
	"{{ .AppName }}/lib/query"
	"{{ .AppName }}/ecode"

	"github.com/gin-gonic/gin"
)

const (
	userID = "uid"
)

const (
	maxLimit = 100
	allLimit = 2000 // 查询所有数据 限制
)

const (
	TypeSuccess = "success"
	TypeError   = "error"
	TypeWarning = "warning"
)

type H map[string]interface{}

type Response struct {
	Code    int         `json:"code"`
	Type    string      `json:"type,omitempty"`// 'success' | 'error' | 'warning';` + `
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

var (
	_succResp      Response = Response{Code: ecode.ErrAllIsOK, Type: TypeSuccess}
	_forbiddenResp Response = Response{Code: ecode.ErrNoPermission, Type: TypeError}
)

func badResp(c *gin.Context, code int, err error, result ...interface{}) {
	resp := Response{
		Code:    code,
		Type:    TypeError,
		Message: err.Error(),
	}
	if len(result) > 0 {
		resp.Result = result[0]
	}
	c.JSON(http.StatusBadRequest, resp)
	c.Abort()

}

func ForbiddenResp(c *gin.Context) {
	c.JSON(http.StatusForbidden, _forbiddenResp)
	c.Abort()
}

func SuccResp(c *gin.Context) {
	c.JSON(http.StatusOK, _succResp)
}

func JsonResp(c *gin.Context, obj interface{}) {
	data := _succResp
	data.Result = obj
	c.JSON(http.StatusOK, data)

}

func getUID(c *gin.Context) string {
	return c.GetString(userID)
}

func paramInt(c *gin.Context, field string) (int64, error) {
	_field, err := strconv.ParseInt(c.Param(field), 10, 64)
	return _field, err
}

func QueryInt64(c *gin.Context, field string) (int64, bool) {
	v, ok := c.GetQuery(field)
	if !ok {
		return 0, ok
	}

	vInt64, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return 0, false
	}
	return vInt64, true
}

func genParams(c *gin.Context) *query.Params {
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 64)
	offset, _ := strconv.ParseInt(c.DefaultQuery("offset", "0"), 10, 64)

	if pageSize, ok := QueryInt64(c, "page_size"); ok {
		limit = pageSize
	}

	if pageSize, ok := QueryInt64(c, "pageSize"); ok {
		limit = pageSize
	}

	if limit > maxLimit {
		limit = maxLimit
	}

	if page, ok := QueryInt64(c, "page"); ok {
		offset = (page - 1) * limit
	}

	orderBy := []query.OrderBy{}
	orderFields := c.QueryArray("orderby")

	for _, field := range orderFields {
		asc := true
		if strings.HasPrefix(field, "-") {
			asc = false
		}
		_field := strings.TrimSpace(strings.TrimPrefix(field, "-"))
		orderBy = append(orderBy, query.OrderBy{Column: _field, Asc: asc})
	}

	if len(orderBy) == 0 {
		orderBy = append(orderBy, query.OrderBy{Column: "id", Asc: false})
	}

	p := &query.Params{
		Limit:   limit,
		Offset:  offset,
		OrderBy: orderBy,
	}

	return p

}

func baseParams() *query.Params {
	p := &query.Params{
		Limit: allLimit, // 查询所有数据限制 防止误用扫全表导致崩溃
	}
	return p
}

func filterParams(c *gin.Context, p *query.Params, fields []string) {
	for _, field := range fields {
		if value, ok := c.GetQuery(field); ok {
			p.Eq(field, value)
		}
	}
}

func filterGteParams(c *gin.Context, p *query.Params, fields []string) {
	for _, field := range fields {
		if value, ok := c.GetQuery(field); ok {
			p.Gte(field, value)
		}
	}
}

func likeParams(c *gin.Context, p *query.Params, fields []string, key string) {
	if value, ok := c.GetQuery(key); ok {
		for _, field := range fields {
			p.Like(field, value)
		}
	}
}