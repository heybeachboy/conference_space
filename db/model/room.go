package model

type Room struct {
	Id          uint64 `gorm:"PRIMARY_KEY;AUTO_INCREMENT;Column:id;comment:主键id" json:"id"`
	SpaceId     uint32 `gorm:"type:varchar(100);Column:root_id;NOT NULL;DEFAULT:'';comment:房间Id" json:"root_id"`
	OwnId       uint32 `gorm:"type:varchar(100);Column:own_id;NOT NULL;DEFAULT:'';" json:"own_id"`
	ChannelName string `gorm:"type:varchar(255);Column:channel_name;NOT NULL;DEFAULT:'';" json:"channel_name"`
	StartAt     int64  `gorm:"type:bigint;Column:start_time;default:0;NOT NULL;" json:"start_at"`
	EndAt       int64  `gorm:"type:bigint;Column:end_time;default:0;NOT NULL;" json:"end_at"`
	UpdatedAt   int64  `gorm:"type:bigint;Column:update_time;default:0;NOT NULL;" json:"updated_at"`
	CreatedAt   int64  `gorm:"type:bigint;Column:create_time;default:0;NOT NULL;" json:"created_at"`
}

func (u Room) TbName() string {
	return "rooms"
}
