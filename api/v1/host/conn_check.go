package host

import (
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "gmcm/model/v1"
	"gmcm/pkg/code"
	"gmcm/pkg/core"
	"gmcm/pkg/utils"
	myssh "gmcm/pkg/utils/ssh"
	"strconv"
	"strings"
)

func ConnectionCheck(c *gin.Context) {
	var hostBindSpec *v1.ConnectionCheckParams
	var failedHost = ""
	var err error
	c.ShouldBindJSON(&hostBindSpec)
	port, _ := strconv.Atoi(hostBindSpec.SSHPort)
	switch hostBindSpec.NoPass {
	case true:
		for _, ip := range hostBindSpec.IP {
			_, err = myssh.NewClient(hostBindSpec.SSHUser, myssh.WithHost(ip), myssh.WithPort(port), myssh.WithAuthByKey(utils.UserHome()+"/.ssh/id_rsa"))
			if err != nil {
				failedHost = failedHost + ip + ","
			}
		}
	case false:
		for _, ip := range hostBindSpec.IP {
			_, err = myssh.NewClient(hostBindSpec.SSHUser, myssh.WithHost(ip), myssh.WithPort(port), myssh.WithPassword(hostBindSpec.SSHPassword), myssh.WithAuthByPass(hostBindSpec.SSHPassword))
			if err != nil {
				failedHost = failedHost + ip + ","
			}
		}
	}

	if failedHost != "" {
		failedHost = strings.TrimRight(failedHost, ",")
		switch hostBindSpec.NoPass {
		case true:
			core.WriteResponse(c, errors.WithCode(code.ErrHostConnectionByKey, "Error connect to remote host: %s, error: %s", failedHost, err.Error()), nil)
			return
		case false:
			core.WriteResponse(c, errors.WithCode(code.ErrHostConnectionByPass, "Error connect to remote host: %s, error: %s", failedHost, err.Error()), nil)
			return
		}
	}

	core.WriteResponse(c, nil, map[string]string{"successes": "ok"})
}
