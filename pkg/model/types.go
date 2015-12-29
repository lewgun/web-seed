package model

import (
	"time"
)

//account type
type Role int
const (
	_ Role = iota
	Normal
	Guest
	Admin
)

//account status
type AccountStatus int
const (
	_ AccountStatus = iota
	Activation
	Inactivation
	Deleted
)

//用户信息
type User struct {
	ID       int    `xorm:" pk autoincr 'id'" json:"id"` //pk
	Uid      string `xorm:" notnull varchar(20) 'uid'" json:"uid"`
	Password string `xorm:"-"`
	Name     string `xorm:" notnull varchar(32) 'name'" json:"name"`
	Phone    string `xorm:" notnull varchar(20) 'phone'" json:"phone"`
	Email    string `xorm:" notnull varchar(64) 'email'" json:"email"`
	Hash     string `xorm:" notnull varchar(60) 'hash'" json:"-"`
	Salt     string `xorm:" notnull varchar(24) 'salt'" json:"-"`

	Role     Role `xorm:" notnull int 'role'" json:"role"`
	Status   AccountStatus `xorm:" notnull int 'status'" json:"status"`
}



type JSONTime string
func (t JSONTime) MarshalJSON() ([]byte, error) {
	const format = "2006-01-02 15:04:05"

	timeNow, _ := time.Parse("2006-01-02T15:04:05+08:00", string(t))
	return []byte(`"` + timeNow.Format(format) + `"`), nil
}
