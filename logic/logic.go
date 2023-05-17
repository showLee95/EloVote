package logic

import (
	"future/algorithm"
	"future/model"
	"future/mysqls"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	p := new(model.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp Err", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数错误",
		})
		return
	}

	intpass, _ := strconv.Atoi(p.Password)
	intrepass, _ := strconv.Atoi(p.RePassword)
	if len(p.Username) == 0 || len(p.Password) == 0 || intpass != intrepass {
		zap.L().Error("SingnUp with invalid param")
		c.JSON(http.StatusOK, gin.H{
			"msg": "请确认密码是否相同",
		})
		return
	}
	if err := mysqls.CheckUserName(p.Username); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户已存在",
		})
	}

	userID := algorithm.GenID()
	user := &model.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysqls.InsertUser(user); err != nil {
		zap.L().Error("mysql insert err ", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})

}

func LoginUp(c *gin.Context) {
	p := new(model.ParamLoginUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("LoginUp Err", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数错误",
		})
		return
	}
	user := &model.User{
		Username: p.Username,
		Password: p.Password,
	}

	//验证密码
	if err := mysqls.LoginUp(user); err != nil {
		zap.L().Error("mysql LoginUp err", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "登录失败，请检查用户名及密码",
		})
		return
	}

	//创造token
	token, err := algorithm.GenerateJWTToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token

	//返回 用户名及token
	c.JSON(http.StatusOK, gin.H{
		"user_name": user.Username,
		"token":     user.Token,
	})
}
