package host

import (
	"bytes"
	v1 "gmcm/model/v1"
	"gmcm/pkg/constand"
	"gmcm/pkg/db"
	"gmcm/pkg/log"
	"gmcm/pkg/utils"
	myssh "gmcm/pkg/utils/ssh"
	"gmcm/static"
	"io/ioutil"
	"strconv"
	"strings"
	"text/template"
)

func refreshHostTemplate() error {
	var hosts []v1.Hosts
	var count int64
	var funcMap = map[string]interface{}{
		"iterate": func(data []v1.Hosts, key string) string {
			var out string
			var tempArr []string
			for _, host := range data {
				switch key {
				case "ip":
					tempArr = append(tempArr, host.IP)
				case "passwd":
					tempArr = append(tempArr, host.SSHPassword)
				}
			}

			tempArr = utils.RemoveRepeatElement(tempArr)
			for _, item := range tempArr {
				out = out + item + `,`
			}

			out = strings.TrimRight(out, ",")
			return out
		},
	}
	var buff bytes.Buffer

	// 当表中没有记录时不需要刷新模板,否则会出错
	db.Client().Table("hosts").Count(&count)
	if count == 0 {
		return nil
	}

	_ = db.Client().Table("hosts").Find(&hosts).Error

	tmpl := template.Must(template.New("ceph_install.conf.tpl").Funcs(funcMap).Parse(string(static.CEPH_INSTALL_CONF_TPL)))
	err := tmpl.Execute(&buff, hosts)
	if err != nil {
		log.Errorf("Error in execute template. Error: %s", err.Error())
		return err
	}

	content, _ := ioutil.ReadAll(&buff)
	err = utils.SaveFile(constand.CEPH_PACKAGE_BASE_PATH+"/ceph_install.conf", string(content))
	if err != nil {
		log.Errorf("error in save template file. Error: %s", err.Error())
		return err
	}
	return nil
}

func newSSHClient(hosts v1.Hosts, ip string) (*myssh.Client, error) {
	var port int
	var cli *myssh.Client
	var err error
	port, _ = strconv.Atoi(hosts.SSHPort)

	switch hosts.NoPass {
	case true:
		cli, err = myssh.NewClient(hosts.SSHUser, myssh.WithAuthByKey(utils.UserHome()+"/.ssh/id_rsa"), myssh.WithHost(ip), myssh.WithPort(port))
	case false:
		cli, err = myssh.NewClient(hosts.SSHUser, myssh.WithPassword(hosts.SSHPassword), myssh.WithHost(ip), myssh.WithPort(port))
	}

	return cli, err
}
