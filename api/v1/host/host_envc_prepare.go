package host

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	v1 "gmcm/model/v1"
	"gmcm/pkg/core"
	"gmcm/static"
)

func EnvCheckPrepare(c *gin.Context) {
	var buf v1.CheckRules

	err := json.Unmarshal(static.CHECK_RULES, &buf)
	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, &buf)
}
