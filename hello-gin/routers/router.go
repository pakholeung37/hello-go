package routers

import (
	"example.com/hello-gin/middleware/jwt"
	"example.com/hello-gin/pkg/setting"
	"example.com/hello-gin/routers/api"
	"example.com/hello-gin/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	var apiV1 = r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/article/:id", v1.GetArticle)
		apiV1.POST("/article", v1.AddArticle)
		apiV1.PUT("/article/:id", v1.EditArticle)
		apiV1.DELETE("/article/:id", v1.DeleteArticle)

	}
	return r
}
