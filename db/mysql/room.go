package mysql

import "ConferenceSpace/db/model"

func NewRoom(r *model.Room) error {
	return dbConn.Table(r.TbName()).Create(r).Error
}
