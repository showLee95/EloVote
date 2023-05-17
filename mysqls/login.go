package mysqls

import (
	"errors"
	"future/model"
)

func LoginUp(user *model.User) (err error) {
	err = CheckUserPass(user)
	if err != nil {
		return err
	}
	return
}

func CheckUserPass(user *model.User) (err error) {
	oPass := user.Password
	password := encyptPassword(oPass)
	sqlstrpass := `select user_id, username, password  from user where username=?`
	if err := db.Get(user, sqlstrpass, user.UserID); err != nil {
		return err
	}
	if password != user.Password {
		return errors.New("用户名或密码错误")
	}
	return
}
