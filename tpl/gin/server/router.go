package server

import (
	h "{{ .AppName }}/handler"
	"{{ .AppName }}/conf"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"{{ .AppName }}/middleware"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title           {{ .AppName }} API
// @version         0.0.1
// @description     This is a sample server celler server.


// @host      localhost:9000
// @BasePath  /api/

// @securityDefinitions.basic  BasicAuth

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used

// @securitydefinitions.oauth2.application  OAuth2Application
// @tokenUrl                                https://example.com/oauth/token
// @scope.write                             Grants write access
// @scope.admin                             Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit  OAuth2Implicit
// @authorizationUrl                     https://example.com/oauth/authorize
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.password  OAuth2Password
// @tokenUrl                             https://example.com/oauth/token
// @scope.read                           Grants read access
// @scope.write                          Grants write access
// @scope.admin                          Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode  OAuth2AccessCode
// @tokenUrl                               https://example.com/oauth/token
// @authorizationUrl                       https://example.com/oauth/authorize
// @scope.admin                            Grants read and write access to administrative information


func InitRouter() *gin.Engine {
	gin.SetMode(conf.Server.Mode)
	gin.DisableConsoleColor()
	r := gin.New()
	r.Use(middleware.GinLogger())
	r.Use(gin.Recovery())

	//health check
	r.GET("/ping", h.PingApi)
	r.GET("/api/ping", h.PingApi)
	r.Use(static.Serve("/", static.LocalFile(conf.Server.StaticDir, true)))

	if conf.Server.Mode != gin.ReleaseMode {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	anonymRoute(r)
	normalRoute(r)

	return r
}

func anonymRoute(r *gin.Engine) {
	r.POST("/api/login", h.Login)
}

func normalRoute(r *gin.Engine) {
	n := r.Group("/api")
	// n.Use(middleware.LoginRequired(conf.Server.JwtSecret))

	// n.GET("/userinfo", h.GetUserInfo)
	// n.GET("/user/permission", h.GetUserMenu)
	n.GET("/logout", h.Logout)

}
