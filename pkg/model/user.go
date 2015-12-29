package model

import (
	"github.com/lewgun/web-seed/pkg/errutil"
)


//FindUser find user by name
func (m *model) FindUser(name string) (*User, error) {
	return nil, errutil.ErrNotFound
}

//DeleteUser delete user by name
func (m *model) DeleteUser(name string) error {
	return nil
}

//UserByUID get the user by uid
func (m *model) UserByUID(uid string) (*User, error) {
	u := &User{}
	has, err := m.Engine.Where("uid = ?", uid).Get(u)
	if !has {
		err = errutil.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return u, nil
}
