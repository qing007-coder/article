package gateway

import (
	"article/pkg/constant"
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine        *gin.Engine
	middleware    *Middleware
	article       *ArticleApi
	administrator *AdministratorApi
	user          *UserApi
}

func NewRouter(base *BaseApi) *Router {
	middleware := NewMiddleware(base)
	administrator := NewAdministratorApi(base)
	article := NewArticleApi(base)
	user := NewUserApi(base)

	return &Router{
		engine:        gin.Default(),
		middleware:    middleware,
		article:       article,
		administrator: administrator,
		user:          user,
	}
}

func (r *Router) Run() {
	r.engine.Use(r.middleware.Cors(), r.middleware.LoggerToFile())
	api := r.engine.Group("api")
	{
		api.POST("reg", r.user.Register)
		api.POST("login", r.user.Login)
		api.POST("sendEmail", r.user.SendVerificationCode)
		api.GET("getUserInfo", r.user.GetUserDetails)
	}

	article := r.engine.Group("article").Use(r.middleware.Auth("", ""))
	{
		article.GET("getArticle", r.article.GetArticleList)
		article.POST("sendArticle", r.article.UploadArticle)
		article.POST("getArticleAllInfo", r.article.GetArticleDetails)
	}

	user := r.engine.Group("user").Use(r.middleware.Auth("", ""))
	{
		user.GET("getUserAllInfo", r.user.GetUserDetails)
		user.POST("modifyPassword", r.user.ModifyPassword)
	}

	privilege := r.engine.Group("privilege")
	{
		privilege.GET("getArticleInQueue", r.middleware.Auth(constant.ARTICLE, constant.JUDEG), r.administrator.GetArticleInQueue)
		privilege.POST("findJudgeArticle", r.middleware.Auth(constant.ARTICLE, constant.JUDEG), r.administrator.GetJudgeArticle)
		privilege.POST("judgeArticles", r.middleware.Auth(constant.ARTICLE, constant.JUDEG), r.administrator.JudgeArticles)
	}

	r.engine.POST("search", r.article.Search).Use(r.middleware.Auth("", ""))
	_ = r.engine.Run(":8080")
}
