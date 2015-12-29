package misc

import (
	"github.com/gin-gonic/gin"
)

func SimpleResponse(c *gin.Context, any interface{}) {

	var obj interface{}
	obj = map[string]string{"result": "success"}

	switch param := any.(type) {
	case error:
		if param != nil {
			c.Error(param)
			obj = gin.H{"result": "fail", "faildesc": param.Error()}
		}

	case map[string]interface{}:
		param["result"] = "success"
		obj = param

	}

	c.JSON(200, obj)
}
