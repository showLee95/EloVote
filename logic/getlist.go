package logic

import (
	"future/mysqls"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ListUp(c *gin.Context) {
	data, err := mysqls.Listdata()
	if err != nil {
		zap.L().Error("List err")
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
	return
}

var combinations [][]string

var vidnames = map[string]int64{}

func Vodedata(c *gin.Context) {
	pairdata := make(map[string]int64)
	for _, keys := range combinations[0] {
		pairdata[keys] = vidnames[keys]
	}
	c.JSON(http.StatusOK, gin.H{
		"data": pairdata,
	})
	return
}
func Vodedata1(c *gin.Context) {

	//获取id 并转换成参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)

	if err != nil {
		zap.L().Error("GET Vode  data err", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": err,
		})
	}
	pairdata := make(map[string]int64)
	for _, keys := range combinations[pid] {
		pairdata[keys] = vidnames[keys]
	}
	c.JSON(http.StatusOK, gin.H{
		"data": pairdata,
	})
	return

}

func VoteUp(c *gin.Context) {
	// vidnames := make(map[string]int64)

	data, err := mysqls.Listdata()
	if err != nil {
		zap.L().Error("VoteUp err")
	}

	for _, dataName := range data {
		vidnames[dataName.Name] = dataName.VId
	}

	combinations = ListCombinations(vidnames)
	c.JSON(http.StatusOK, gin.H{
		"data": combinations,
	})
	return

}

func ListCombinations(vidnames map[string]int64) [][]string {
	combinations := [][]string{}

	// 获取所有键值对的键
	keys := make([]string, 0, len(vidnames))
	for key := range vidnames {
		keys = append(keys, key)
	}

	// 嵌套循环遍历所有组合方式
	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			combination := []string{keys[i], keys[j]}
			combinations = append(combinations, combination)
		}
	}

	return combinations
}
