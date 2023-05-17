package router

import (
	"future/algorithm"
	"future/loger"
	"future/logic"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(loger.GinLogger(), loger.GinRecovery(true))
	r.POST("/login", logic.LoginUp)
	r.POST("/sign", logic.SignUp)
	v1 := r.Use(algorithm.JWTAuthMiddleware())
	{
		v1.GET("/list", logic.ListUp)
		v1.GET("/list2", logic.VoteUp)
		v1.GET("/list3/:id", logic.Vodedata1)
		v1.POST("/make", logic.MakeVote)
		v1.POST("/vote", logic.VoteServer)
	}
	r.GET("/ping", algorithm.JWTAuthMiddleware(), func(c *gin.Context) {

		c.String(http.StatusOK, "pong")

	})
	return r
}
