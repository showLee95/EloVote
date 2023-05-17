package mysqls

import (
	"errors"
	"future/model"

	"go.uber.org/zap"
)

func Listdata() (votelist []*model.Vote, err error) {
	sqlStr := "select vid, votename, vote, rate, score from vote"
	if err := db.Select(&votelist, sqlStr); err != nil {
		zap.L().Error("Sql select err", zap.Error(err))
		err = nil
	}
	return
}

// insert into vote (vid, votename, vote, rate, score)VALUES (1212,"key2",0,0.55,21)

func CreateVoteData(p *model.Vote) (err error) {
	p.Vote = 0
	p.Rate = 0
	p.Score = 1400
	sqlStr := `insert into vote (vid, votename, vote, rate, score)
VALUES  (?,?,?,?,?)
`
	_, err = db.Exec(sqlStr, p.VId, p.Name, p.Vote, p.Rate, p.Score)
	return
}

func CheckVoteName(name string) (err error) {
	sqlStr := `select count(id) from vote where votename=?`
	var count int
	if err := db.Get(&count, sqlStr, name); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}
