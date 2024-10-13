package model

type LoginReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Account          string `json:"account" binding:"min=3,max=9"`
	Password         string `json:"password" binding:"min=8,max=20"`
	Email            string `json:"email" binding:"email"`
	VerificationCode string `json:"verification_code" binding:"min=6,max=6"`
}

type SendVerificationCode struct {
	Email string `json:"email"`
}