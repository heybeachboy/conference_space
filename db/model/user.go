package model

import "ConferenceSpace/constant"

type User struct {
	Uid       uint32             `gorm:"PRIMARY_KEY;AUTO_INCREMENT;Column:uid;" json:"uid"`
	Username  string             `gorm:"type:varchar(100);Column:username;NOT NULL;DEFAULT:'';" json:"username"`
	Password  string             `gorm:"type:varchar(100);Column:password;NOT NULL;DEFAULT:'';" json:"password"`
	Nickname  string             `gorm:"type:varchar(100);Column:nickname;NOT NULL;DEFAULT:'';" json:"nickname"`
	Avatar    string             `gorm:"type:varchar(100);Column:avatar;NOT NULL;DEFAULT:'';" json:"avatar"`
	Phone     string             `gorm:"type:varchar(20);Column:phone;NOT NULL;DEFAULT:'';" json:"phone"`
	Email     string             `gorm:"type:varchar(50);Column:email;NOT NULL;DEFAULT:'';" json:"email"`
	Gender    constant.GenderTyp `gorm:"type:uint;Column:gender;NOT NULL;DEFAULT:0;" json:"gender"`
	UpdatedAt int64              `gorm:"type:bigint;Column:update_time;default:0;NOT NULL;" json:"updated_at"`
	CreatedAt int64              `gorm:"type:bigint;Column:create_time;default:0;NOT NULL;" json:"created_at"`
}

func (u User) TbName() string {
	return "users"
}
