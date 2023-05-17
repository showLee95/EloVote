package logic

import (
	"future/model"
	"future/mysqls"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func VoteServer(c *gin.Context) {
	p := new(model.Votename)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("error", zap.Any("err", err))
		zap.L().Error("Create vote param")
		return
	}

	err := mysqls.Lookup(p)
	if err != nil {
		zap.L().Error("数据库写入错误")
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "投票成功",
	})
	return
}
