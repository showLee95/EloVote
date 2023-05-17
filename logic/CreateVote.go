package logic

import (
	"future/algorithm"
	"future/model"
	"future/mysqls"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func MakeVote(c *gin.Context) {
	p := new(model.Vote)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("error", zap.Any("err", err))
		zap.L().Error("Create vote param")
		return
	}
	//检查是否重复提交name
	err := mysqls.CheckVoteName(p.Name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "该名字已存在",
		})
		return
	}

	//生成随机ID
	p.VId = algorithm.GenID()
	//添加到数据库
	err = mysqls.CreateVoteData(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "添加失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "添加成功",
	})
	return
}
