package host

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "gmcm/model/v1"
	"gmcm/pkg/code"
	"gmcm/pkg/core"
	"gmcm/pkg/utils"
	myssh "gmcm/pkg/utils/ssh"
	"gmcm/static"
	"strings"
	"time"
)

func EnvcAction(c *gin.Context) {
	var hosts v1.Hosts
	var buf v1.CheckRules
	var cli *myssh.Client
	var out v1.ScriptOut

	err := c.ShouldBindJSON(&hosts)
	if err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "Error bind request body"), nil)
		return
	}

	// 写入脚本文件到本地
	if utils.IsNotExist("/tmp/check.sh") {
		_ = utils.SaveFile("/tmp/check.sh", string(static.CHECK_SCRIPTS))
	}

	// 解析rules
	_ = json.Unmarshal(static.CHECK_RULES, &buf)

	ips := strings.Split(hosts.IP, ",")
	for _, ip := range ips {
		cli, _ := newSSHClient(hosts, ip)
		_ = cli.UploadFile("/tmp/check.sh", "/tmp/env_check.sh")
	}

	clone := buf
	var start = time.Now()
	for idx, rule := range buf.Rules {
		var messages = []string{}
		var checkStatus = "OK"

		funcStart := time.Now()

		for _, ip := range ips {
			sshStart := time.Now()
			cli, err = newSSHClient(hosts, ip)
			if err != nil {
				core.WriteResponse(c, errors.WithCode(code.ErrHostConnectionByPass, "无法与目标主机建立链接.Error: %s", err.Error()), nil)
				return
			}
			sshCost := time.Since(sshStart).Seconds()
			fmt.Printf("IP: %s ssh connection cost: %f\n", ip, sshCost)

			cmd := `bash /tmp/env_check.sh ` + rule.Func

			o, _ := cli.Output(cmd)
			_ = json.Unmarshal(o, &out)

			if out.Status == "FAILED" {
				checkStatus = "FAILED"
				messages = append(messages, ip+`: `+out.Message)
			}
		}
		message := utils.ToString(messages, ",")
		clone.Rules[idx].Message = message
		clone.Rules[idx].Status = checkStatus
		funcCost := time.Since(funcStart).Seconds()
		fmt.Printf("func: %s total cost: %f\n", rule.Func, funcCost)
	}
	var cost = time.Since(start).Seconds()
	fmt.Printf("env check total cost: %f\n", cost)

	core.WriteResponse(c, nil, &clone)
}
