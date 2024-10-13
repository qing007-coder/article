package gateway

import (
	"article/pkg/errors"
	"article/pkg/model"
	"article/pkg/tools"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserApi struct {
	*BaseApi
}

func NewUserApi(base *BaseApi) *UserApi {
	return &UserApi{
		base,
	}
}

func (u *UserApi) Login(ctx *gin.Context) {
	var req model.LoginReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	var user model.User
	if err := u.db.Where("account = ?", req.Account).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			tools.BadRequest(ctx, errors.RecordNotFound.Error())
			return
		} else {
			tools.BadRequest(ctx, errors.OtherError.Error())
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		if err != bcrypt.ErrMismatchedHashAndPassword {
			tools.BadRequest(ctx, "wrong password")
			return
		} else {
			tools.BadRequest(ctx, errors.OtherError.Error())
			return
		}
	}

	tools.StatusOK(ctx, gin.H{
		"token": tools.CreateToken(user.ID, u.conf),
	}, "登录成功")
}

func (u *UserApi) Register(ctx *gin.Context) {
	var req model.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	var count int64
	u.db.Where("email = ?", req.Email).Count(&count)
	if count > 0 {
		tools.BadRequest(ctx, "邮箱已注册")
		return
	}

	u.db.Where("account = ?", req.Account).Count(&count)
	if count > 0 {
		tools.BadRequest(ctx, "账户已存在")
		return
	}

	data, err := u.rdb.Get(u.ctx, req.Email)
	if err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}
	if data != req.VerificationCode {
		tools.BadRequest(ctx, "验证码错误")
		return
	}

	id := tools.CreateID()
	u.db.Create(&model.User{
		ID:       id,
		Account:  req.Account,
		Password: req.Password,
		Email:    req.Email,
		Status:   true,
	})

	token, err := tools.CreateToken(id, u.conf)
	if err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	tools.StatusOK(ctx, gin.H{
		"token": token,
	}, "注册成功")
}

func (u *UserApi) SendVerificationCode(ctx *gin.Context) {
	var req model.SendVerificationCodeReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		tools.BadRequest(ctx, err.Error())
		return
	}

	data, err := u.rdb.Get(u.ctx, req.Email+".send")
	if err == nil && data != "" {
		tools.BadRequest(ctx, "发送验证码过于频繁")
		return
	}

	code, err := u.email.SendVerificationCode(req.Email)
	if err != nil {
		tools.BadRequest(ctx, "验证码发送失败")
		return
	}

	if err := u.rdb.Set(u.ctx, req.Email+".send", 1, time.Minute); err != nil {
		tools.BadRequest(ctx, "发送失败")
		return
	}

	if err := u.rdb.Set(u.ctx, req.Email, code, time.Minute*10); err != nil {
		tools.BadRequest(ctx, "发送失败")
		return
	}

	tools.StatusOK(ctx, nil, "发送成功")
}

func (u *UserApi) GetUserDetails(ctx *gin.Context) {

}

func (u *UserApi) ModifyPassword(ctx *gin.Context) {

}
