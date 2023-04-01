package controller

import (
	"elderly/dao/mysql"
	"elderly/logic"
	"elderly/models"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 注册接口
// @Summary 用户注册接口
// @Description 用户注册接口
// @Tags 用户注册接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamSignUp true "注册参数"
// @Success 200 {object} _ResponseSignUpList
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	//1.获取参数,参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) // 翻译错误为中文并返回
		return
	}
	//2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp() failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登录接口
// @Summary 用户登录接口
// @Description 用户登录接口
// @Tags 用户登录接口
// @Accept application/json
// @Produce application/json
// @Param object body models.ParamLogin true "登录参数"
// @Success 200 {object} _ResponseLoginList
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	//1.获取参数,校验参数
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误，直接返回响应
		zap.L().Error("login with invalid param", zap.Error(err))
		//判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans))) // 翻译错误为中文并返回
		return
	}
	//2.业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login() failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
