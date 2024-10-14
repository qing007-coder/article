package gateway

import (
	"article/pkg/model"
	"article/pkg/tools"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type AdministratorApi struct {
	*BaseApi
}

func NewAdministratorApi(base *BaseApi) *AdministratorApi {
	return &AdministratorApi{
		base,
	}
}

func (a *AdministratorApi) GetArticleInQueue(ctx *gin.Context) {
	var articleJudgeRecords []model.ArticleJudgeRecord
	result := a.db.Where("flag = ?", false).Find(&articleJudgeRecords)
	if result.Error != nil {
		tools.BadRequest(ctx, result.Error.Error())
		return
	}

	var queue []model.Queue
	for i := 0; i < len(articleJudgeRecords); i++ {
		queue = append(queue, model.Queue{
			ArticleId: articleJudgeRecords[i].ArticleID,
		})
	}

	tools.StatusOK(ctx, gin.H{
		"queue": queue,
	}, "提取成功")
}

func (a *AdministratorApi) GetJudgeArticle(ctx *gin.Context) {
	var req model.GetJudgeArticleReq
	if err := ctx.ShouldBind(&req); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	data, err := a.es.GetDocumentByID(req.ArticleId)
	if err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	var article model.Article
	if err := json.Unmarshal(data, &article); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}
	tools.StatusOK(ctx, gin.H{
		"articleInfo": article,
	}, "查找成功")
}
