package mysqls

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"future/model"

	"go.uber.org/zap"
)

const secret = "lee"

func CheckUserName(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已存在")
	}
	return
}

func InsertUser(user *model.User) (err error) {
	//对密码加密
	user.Password = encyptPassword(user.Password)
	//SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	if err != nil {
		zap.L().Error("MYSQL INSER ERR :", zap.Error(err))
	}
	return
}

func encyptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Sum([]byte(oPassword))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
