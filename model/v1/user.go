package v1

import (
	"github.com/google/uuid"
	"gmcm/pkg/utils/auth"
)

type User struct {
	ObjectMeta  `json:"metadata,omitempty"`
	UUID        uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	UserName    string    `json:"username" gorm:"column:name;comment:用户登录名"`
	Password    string    `json:"-" gorm:"column:password;comment:用户登录密码"`
	NickName    string    `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`
	HeaderImg   string    `json:"headerImg" gorm:"comment:用户头像"`
	AuthorityId string    `json:"authorityId" gorm:"default:888;comment:用户角色ID"`
}

func (u *User) Compare(pwd string) (err error) {
	err = auth.Compare(u.Password, pwd)

	return
}
