package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	UUID     int64  `gorm:"column:id; type:int; index:idx_uuid; primaryKey; autoIncrement; comment: 'primary key is ID"`
	RoleName string `gorm:"column:role_name"`
	RoleNote string `gorm:"column:role_note; type:text"`
}

func (r *Role) TableName() string {
	return "go_db_user"
}
