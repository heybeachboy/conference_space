package mysql

import "ConferenceSpace/db/model"

func GetUserByUsername(name string) *model.User {
	u := new(model.User)
	dbConn.Table(u.TbName()).Where("`username` = ?", name).First(u)
	return u
}

func CreateUser(u *model.User) error {
	return dbConn.Table(u.TbName()).Create(u).Error
}
