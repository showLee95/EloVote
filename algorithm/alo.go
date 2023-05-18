package algorithm

import (
	"math"
)

const (
	kFactor = 32 // K-factor for Elo rating system
)

// CalculateElo calculates the new Elo ratings for two players after a match
func CalculateElo(p_Victory int64, p_Burden int64) (a int64, err error) {
	var exponent float64
	diff := p_Victory - p_Burden
	exponent = float64(diff) / 400
	probability := 1 / (1 + math.Pow(10, exponent))
	victory := 32 * (1 - probability)
	a = int64(victory)
	// zap.L().Error("p_Victory, sqlStr, Victory", zap.Error(errors.New(a)))
	return a, nil

	// A队1500分，B队1600分，则
	// 预估A队的胜负值 Ea = 1/(1+10^[(1600-1500)/400]) = 0.36
	// 预估B队的胜负值 Eb = 1/(1+10^[(1500-1600)/400]) = 0.64

	// 假设A队赢了，
	// A队最终得分为 R'a = 1500 + 32*(1-0.36) = 1500+20.5 = 1520, 赢20分，B队输20分。
	// 假设B队赢了，
	// B队最终得分为 R'b = 1600 + 32*(1-0.64) = 1600 + 11.52 = 1612, 赢12分，A队输12分。

	////
	////
	////

	// diff := rating2 - rating1

	// // 计算指数
	// exponent := diff / constant

	// // 计算胜率
	// probability := 1 / (1 + math.Pow(10, exponent))
}
