package controller

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-deck/app/controller/dto"
	"go-deck/app/global"
	"go-deck/app/model/entity"
	"go-deck/app/response"
	"go-deck/app/service"
	myJwt "go-deck/pkg/jwt"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body dto.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func Login(c *gin.Context) {
	var req dto.Login

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage("参数验证失败", c)
		c.Abort()
		return
	}

	// 校验验证码
	//if !store.Verify(req.CaptchaId, req.Captcha, true) {
	//	response.FailWithMessage("验证码错误", c)
	//	c.Abort()
	//	return
	//}

	// 查询用户
	user, err := global.DB.FindUserByUserName(req.Username)
	if err != nil {
		response.FailWithMessage("用户名不存在", c)
		c.Abort()
		return
	}

	// 校验密码
	equal := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if equal != nil {
		response.FailWithMessage("用户名不存在或者密码错误", c)
		c.Abort()
		return
	}

	// 生成token
	tokenNext(c, *user)
}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user entity.User) {
	j := myJwt.NewJWT(&myJwt.Config{SigningKey: global.Config.JwtConfig.SigningKey})
	claims := myJwt.CustomClaims{
		ID:         user.ID,
		NickName:   user.NickName,
		Username:   user.Username,
		BufferTime: 60 * 60 * 24, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,       // 签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*7, // 过期时间 7天
			Issuer:    "deck",                         // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.Logger.Error("获取token失败", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	response.OkWithDetailed(response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
}

// @Tags User
// @Summary 创建用户
// @Produce  application/json
// @Param data body dto.CreateUser true "用户名, 昵称, 密码, 角色ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/create [post]
func CreateUser(c *gin.Context) {
	var req dto.CreateUser

	// 校验请求参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage("参数验证失败", c)
		c.Abort()
		return
	}

	// 判断用户名是否存在
	u, err := global.DB.FindUserByUserName(req.Username)
	if err == nil && u != nil {
		response.FailWithMessage("该用户名已注册", c)
		c.Abort()
		return
	}

	// 创建用户
	user := entity.User{
		Username: req.Username,
		Password: req.Password,
	}
	err = global.DB.CreateUser(&user)
	if err != nil {
		global.Logger.Error("注册失败", zap.Any("err", err))
		response.FailWithMessage("注册失败", c)
	} else {
		response.OkWithMessage("注册成功", c)
	}
}

// 给用户绑定角色权限
func BindRole(c *gin.Context) {
	var req dto.BindUserRole

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage("参数验证失败", c)
		c.Abort()
		return
	}

	// 绑定角色
	if err := service.BindUserRole(req.UserId, req.RoleIds); err != nil {
		response.FailWithMessage("绑定角色失败", c)
		c.Abort()
		return
	}

	response.OkWithMessage("绑定角色成功", c)
	return
}
