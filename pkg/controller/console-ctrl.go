package controller
import (

	"github.com/lewgun/web-seed/pkg/controller/internal"
	"github.com/lewgun/web-seed/pkg/misc"

	"github.com/gin-gonic/gin"
)


// logout endpoint
func Logout(ctx *gin.Context) {

	internal.DelSession(ctx)

	m := map[string]interface{}{}
	misc.SimpleResponse(ctx, m)

}