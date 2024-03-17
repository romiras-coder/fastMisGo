package model

type AuthenticationInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type UserResp struct {
	UserId   int    `json:"userId" binding:"required"`
	UserName string `json:"userName" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
