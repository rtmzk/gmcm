package host

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hpcloud/tail"
	"github.com/marmotedu/errors"
	v1 "gmcm/model/v1"
	"gmcm/pkg/code"
	"gmcm/pkg/core"
	"gmcm/pkg/log"
	"io"
	"strconv"
)

func GetInstallLogs(c *gin.Context) {
	log.Debug("get install logs api called.")

	var logs v1.InstallLogs

	param := c.Param("offset")
	offset, _ := strconv.Atoi(param)
	logs.Offset = int64(offset)
	logs.Text = ""

	tailConf := tail.Config{
		Location: &tail.SeekInfo{
			Offset: logs.Offset,
			Whence: io.SeekStart,
		},
		Follow:    false,
		MustExist: true,
		Logger:    tail.DiscardingLogger,
	}

	t, err := tail.TailFile("/tmp/ceph_install.log", tailConf)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrLogFileNotFound, "未找到安装日志，请先初始化存储, error: %s", err.Error()), nil)
		return
	}

	for line := range t.Lines {
		logs.Offset = logs.Offset + int64(len(line.Text)) + 1
		logs.Text = fmt.Sprintf(logs.Text + line.Text + "\n")
	}

	core.WriteResponse(c, nil, &logs)
}
