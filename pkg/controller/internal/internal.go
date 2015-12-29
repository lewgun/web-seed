package internal

import (
	"github.com/lewgun/web-seed/pkg/model"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const sessKey = "account"

func DelSession(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sess.Delete(sessKey)
	sess.SaveDelete()
}

func SaveSession(ctx *gin.Context, uid string) {
	sess := sessions.Default(ctx)
	sess.Set(sessKey, uid)
	sess.Save()
}

func Session(ctx *gin.Context) interface{} {
	s := sessions.Default(ctx)
	return s.Get(sessKey)

}

//IsAuthorized whether authorized or not
func IsAuthorized(c *gin.Context) bool {

	if data := Session(c); data != nil {
		return true
	}

	return false
}

//CurrentUser get the current user.
func CurrentUser(c *gin.Context) *model.User {

	var (
		uid string
		ok  bool
	)

	data := Session(c)

	if data == nil {
		return nil
	}

	if uid, ok = data.(string); !ok {
		return nil
	}

	u, err := model.M.UserByUID(uid)
	if err != nil {
		return nil
	}

	return u

}
