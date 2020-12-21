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
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	// 校验验证码
	//if !store.Verify(req.CaptchaId, req.Captcha, true) {
	//	response.FailWithMessage("验证码错误")
	//	c.Abort()
	//	return
	//}

	// 查询用户
	user, err := global.DB.FindUserByUserName(req.Username)
	if err != nil {
		response.FailWithMessage(c, "用户名不存在")
		c.Abort()
		return
	}

	// 校验密码
	equal := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if equal != nil {
		response.FailWithMessage(c, "用户名不存在或者密码错误")
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
		response.FailWithMessage(c, "获取token失败")
		return
	}

	response.OkWithDetailed(c, response.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功")
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
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	// 判断用户名是否存在
	u, err := global.DB.FindUserByUserName(req.Username)
	if err == nil && u != nil {
		response.FailWithMessage(c, "该用户名已注册")
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
		response.FailWithMessage(c, "注册失败")
	} else {
		response.OkWithMessage(c, "注册成功")
	}
}

// @Tags User
// @Summary 删除User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DeleteUser true "ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/delete [post]
func DeleteUser(c *gin.Context) {
	var req dto.DeleteUser
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if err := service.DeleteByIds(entity.User{}, req.Ids); err != nil {
		global.Logger.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage(c, "删除失败")
	} else {
		response.OkWithMessage(c, "删除成功")
	}
}

// @Tags User
// @Summary 更新User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.UpdateUser true "用户名"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /user/update [post]
func UpdateUser(c *gin.Context) {
	var req dto.UpdateUser

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	api := entity.User{
		ID:        req.Id,
		Username:  req.Username,
		Password:  req.Password,
	}

	if err := service.UpdateById(&api, &entity.User{}, req.Id); err != nil {
		global.Logger.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage(c, "修改失败")
	} else {
		response.OkWithMessage(c, "修改成功")
	}
}

// @Tags User
// @Summary 分页获取User列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.SearchUser true "分页获取User列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/list [post]
func GetUserList(c *gin.Context) {
	var req dto.SearchUser
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	if total, list, err := service.GetUserList(req); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithDetailed(c, response.PageResult{
			List:  list,
			Total: total,
			Page:  req.Page,
			Size:  req.Size,
		}, "获取成功")
	}
}

// @Tags User
// @Summary 根据id获取User
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.DetailUser true "根据id获取User"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/detail [post]
func GetUserById(c *gin.Context) {
	var req dto.DetailUser
	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	role := entity.User{}
	err := service.DetailById(&role, req.Id)
	if err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithData(c, role)
	}
}

// @Tags User
// @Summary 获取所有的User 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/all [post]
func GetAllUsers(c *gin.Context) {
	if apis, err := global.DB.AllUsers(); err != nil {
		global.Logger.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage(c, "获取失败")
	} else {
		response.OkWithDetailed(c, apis, "获取成功")
	}
}


// @Tags User
// @Summary 给用户绑定角色权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dto.BindUserRole true "用户id, 角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"绑定成功"}"
// @Router /user/bindRole [post]
func BindRole(c *gin.Context) {
	var req dto.BindUserRole

	// 校验参数
	if err := c.ShouldBind(&req); err != nil {
		response.FailWithMessage(c, "参数验证失败")
		c.Abort()
		return
	}

	// 绑定角色
	if err := service.BindUserRole(req.UserId, req.RoleIds); err != nil {
		response.FailWithMessage(c, "绑定角色失败")
		c.Abort()
		return
	}

	response.OkWithMessage(c, "绑定角色成功")
	return
}
