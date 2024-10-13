package gateway

import "github.com/gin-gonic/gin"

type UserApi struct {
	*BaseApi
}

func NewUserApi(base *BaseApi) *UserApi {
	return &UserApi{
		base,
	}
}

func (u *UserApi) Login(ctx *gin.Context) {

}

func (u *UserApi) Register(ctx *gin.Context) {

}

func (u *UserApi) SendVerificationCode(ctx *gin.Context) {

}

func (u *UserApi) GetUserDetails(ctx *gin.Context) {

}

func (u *UserApi) ModifyPassword(ctx *gin.Context) {

}
