package schema

import (
	"{{ .AppName }}/ecode"
	"net/http"
	"reflect"

	"github.com/go-playground/locales/zh_Hans_CN"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var validate *validator.Validate
var trans ut.Translator

func init() {
	trans, _ = ut.New(zh_Hans_CN.New()).GetTranslator("zh")
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("label")
	})
	if err := zh.RegisterDefaultTranslations(validate, trans); err != nil {
		zap.L().Error("register validator failed", zap.Error(err))
	}
}

func Validator(data interface{}, c *gin.Context, b binding.Binding) (err error) {
	if err = c.ShouldBindWith(data, b); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: ecode.ErrInvalidParam, Message: err.Error()})
		return err
	}
	if err = validate.Struct(data); err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: ecode.ErrInvalidParam, Message: err.(validator.ValidationErrors)[0].Translate(trans)})
		return err
	}
	return nil
}

type Response struct {
	Code    int         `json:"code"`
	Type    string      `json:"type,omitempty"` // 'success' | 'error' | 'warning';
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

type Login struct {
	UID      string `form:"uid" json:"uid" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
