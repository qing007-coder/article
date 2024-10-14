package gateway

import (
	"article/pkg/model"
	"article/pkg/tools"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
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

func (a *AdministratorApi) JudgeArticles(ctx *gin.Context) {
	var req model.JudgeArticleReq
	if err := ctx.ShouldBind(&req); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	if err := a.es.Update(req.ArticleID, map[string]interface{}{
		"status": "1",
	}); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	a.db.Where("article_id = ?", req.ArticleID).Updates(&model.ArticleJudgeRecord{
		IsJudge:         true,
		AdministratorID: ctx.GetString("user_id"),
		Result:          req.Status,
		JudgeTime:       time.Now(),
	})

	tools.StatusOK(ctx, nil, "审批成功")
}
