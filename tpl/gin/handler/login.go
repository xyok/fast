package handler

import (
	"{{ .AppName }}/schema"
	e "{{ .AppName }}/ecode"
	"github.com/gin-gonic/gin"
)

// func GenToken(m *model.User) string {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"uid":   m.UID,
// 		"id":    m.ID,
// 		"name":  m.Name,
// 		"super": m.IsSuper,
// 		"exp":   time.Now().Add(time.Hour * 72).Unix(),
// 	})
// 	tokenStr, _ := token.SignedString([]byte(conf.Server.JwtSecret))
// 	return tokenStr
// }

// @Summary login
// @Description 登录
// @Tags auth
// @Accept application/json
// @Produce application/json
// @param param body schema.Login true "请求参数,字段说明点击model"
// @Success 200 {object} schema.Response
// @Failure 400 {object} schema.Response
// @Router /login [post]
func Login(c *gin.Context) {
	var msg schema.Login
	if err := c.ShouldBind(&msg); err != nil {
		badResp(c, e.ErrInvalidParam, err)
		return
	}

	// var user model.User
	// code, err := s.User.Login(&msg, &user)
	// if err != nil {
	// 	badResp(c, code, err)
	// 	return
	// }

	// token := GenToken(&user)

	JsonResp(c, H{"token": "token"})
}

func Logout(c *gin.Context) {
	//todo logout log and clear token
	SuccResp(c)
}