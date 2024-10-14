package gateway

import (
	"article/pkg/errors"
	"article/pkg/model"
	"article/pkg/tools"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type ArticleApi struct {
	*BaseApi
}

func NewArticleApi(base *BaseApi) *ArticleApi {
	return &ArticleApi{
		base,
	}
}

func (a *ArticleApi) UploadArticle(ctx *gin.Context) {
	var req model.UploadArticleReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	uid := ctx.GetString("user_id")
	id := tools.CreateID()
	if err := a.es.CreateDocument(&model.Article{
		ID:       id,
		AuthorID: uid,
		Time:     time.Now(),
		Read:     0,
		Like:     0,
		Title:    req.Title,
		Content:  req.Content,
		Status:   "0",
	}, id); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	a.db.Create(&model.ArticleJudgeRecord{
		ArticleID: id,
		IsJudge:   false,
		JudgeTime: time.Now(),
	})

	tools.StatusOK(ctx, nil, "发布成功，待审批")
}

func (a *ArticleApi) Search(ctx *gin.Context) {
	var req model.SearchReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	shouldQueries := []types.Query{
		{
			Match: map[string]types.MatchQuery{
				"title": {
					Query:     req.Content,
					Fuzziness: "AUTO",
				},
			},
		},
		{
			Match: map[string]types.MatchQuery{
				"content": {
					Query:     req.Content,
					Fuzziness: "AUTO",
				},
			},
		},
	}

	mustQueries := []types.Query{
		{
			MatchPhrase: map[string]types.MatchPhraseQuery{
				"status": {Query: "1"},
			},
		},
	}
	data, err := a.es.Search(mustQueries, shouldQueries, nil, 0, 10)
	if err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	var list []model.Article
	for _, v := range data.Hits.Hits {
		var article model.Article
		if err := json.Unmarshal(v.Source_, &article); err != nil {
			tools.BadRequest(ctx, err.Error())
			return
		}

		list = append(list, article)
	}

	tools.StatusOK(ctx, gin.H{
		"data": list,
	}, "搜索成功")
}

func (a *ArticleApi) GetArticleList(ctx *gin.Context) {
	data, err := a.es.GetList(10, 0)
	if err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	var list []model.Article
	for _, v := range data.Hits.Hits {
		var article model.Article
		if err := json.Unmarshal(v.Source_, &article); err != nil {
			tools.BadRequest(ctx, err.Error())
			return
		}
		list = append(list, article)
	}

	tools.StatusOK(ctx, gin.H{
		"data": list,
	}, "获取成功")
}

func (a *ArticleApi) GetArticleDetails(ctx *gin.Context) {
	var req model.GetArticleDetailsReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	data, err := a.es.GetDocumentByID(req.ID)
	if err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	var article model.Article
	if err := json.Unmarshal(data, &article); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	var user model.User
	if err := a.db.Where("id = ?", article.AuthorID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			tools.BadRequest(ctx, errors.RecordNotFound.Error())
			return
		} else {
			tools.BadRequest(ctx, errors.OtherError.Error())
			return
		}
	}

	tools.StatusOK(ctx, gin.H{
		"author":  user,
		"article": article,
	}, "查询成功")
}
