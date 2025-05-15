package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     uuid.UUID `gorm:"column:uuid; type:varchar(255); index:idx_uuid; not null; unique"`
	Username string    `gorm:"column:username"`
	IsActive bool      `gorm:"column:is_active; type:boolean;"`
	Roles    []Role    `gorm:"many2many:user_roles;"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
