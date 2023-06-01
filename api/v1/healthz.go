package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gmcm/pkg/utils/version"
)

func GetHealthz(c *gin.Context) {
	data := map[string]string{
		"status":      "available",
		"environment": gin.Mode(),
		"version":     version.GitVersion,
	}
	js, _ := json.MarshalIndent(&data, "", "    ")
	_, _ = c.Writer.Write(js)
}
