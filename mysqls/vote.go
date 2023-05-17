package mysqls

import (
	"future/algorithm"
	"future/model"

	"go.uber.org/zap"
)

func Lookup(p *model.Votename) (err error) {
	Victory := p.Victory
	Burden := p.Burden
	p_Victory := new(model.Vote)
	p_Burden := new(model.Vote)
	sqlStr := `select vid, votename, vote, rate, score from vote where votename=?`
	if err := db.Get(p_Victory, sqlStr, Victory); err != nil {
		zap.L().Error("p_Victory, sqlStr, Victory")
		return err
	}
	if err := db.Get(p_Burden, sqlStr, Burden); err != nil {
		zap.L().Error("p_Burden, sqlStr, Burdeny")
		return err
	}
	p1 := p_Victory.Score

	p2 := p_Burden.Score
	EloScore, err := algorithm.CalculateElo(p1, p2)
	if err != nil {
		zap.L().Error("计算失败")
		return
	}
	VictoryScore := p_Victory.Score + EloScore
	BurdenScore := p_Burden.Score - EloScore
	sqlEloScore := "UPDATE  vote SET  score= ? WHERE votename = ?  "
	_, err = db.Exec(sqlEloScore, VictoryScore, Victory)
	if err != nil {
		zap.L().Error("修改失败")
		return
	}
	_, err = db.Exec(sqlEloScore, BurdenScore, Burden)
	if err != nil {
		zap.L().Error("修改失败")
		return
	}
	return
}
