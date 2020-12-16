package dto

// User login structure
type Login struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
	CaptchaId string `json:"captchaId" binding:"required"`
}

// User register structure
type CreateUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//
type BindUserRole struct {
	UserId  int   `json:"userId" binding:"required"`
	RoleIds []int `json:"roleIds" binding:"required"`
}