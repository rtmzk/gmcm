package host

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	v1 "gmcm/model/v1"
	"gmcm/pkg/code"
	"gmcm/pkg/constand"
	"gmcm/pkg/core"
	"gmcm/pkg/db"
	"gmcm/pkg/log"
	"gmcm/pkg/utils"
	"gmcm/pkg/utils/devices"
	"gmcm/pkg/utils/storage"
	mytemplate "gmcm/pkg/utils/template"
	"gmcm/static"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"
)

func HostInit(c *gin.Context) {
	var devs v1.DeviceList
	var hosts []v1.Hosts
	var devsmap = make(map[string][]devices.Devices)
	var flushTemplate v1.FlushTemplate
	var funcMap = mytemplate.NewFuncMap()
	var totalCount int
	var indexCount = 0
	var errChan = make(chan error)
	var doneChan = make(chan interface{})

	defer close(errChan)
	defer close(doneChan)

	if err := c.ShouldBindJSON(&devs); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, "Bad Request"), nil)
		return
	}

	for _, item := range devs.Devices {
		for _, dev := range item.Device {
			// 过滤掉禁用的设备
			if dev.Enabled == true {
				devsmap[item.IP] = append(devsmap[item.IP], dev)
			}

			// 设置磁盘类型,如果和系统识别的默认类型不符则远程执行命令修改
			if dev.Type != dev.DefaultType && dev.Type == "ssd" {
				var host v1.Hosts
				db.Client().Table("hosts").Where("ip = ?", item.IP).First(&host)

				cli, _ := newSSHClient(host, item.IP)
				cmd := `echo 0 > /sys/block/` + dev.Name + `/queue/rotational`
				cli.Run(cmd)
			}

			// 获取索引盘数量
			if dev.Enabled && dev.Type == "ssd" && !dev.Cached {
				indexCount = indexCount + 1
			}
		}
	}
	// 统计磁盘总数
	for i := 0; i < len(devs.Devices); i++ {
		totalCount += len(devsmap[devs.Devices[i].IP])
	}

	// 计算pg总数
	replicas := devs.Replicas
	poolPgs := totalCount * 100 / replicas / 1
	indexPg := indexCount * 100 / replicas / 1

	//获取节点列表
	db.Client().Table("hosts").Find(&hosts)

	//计算data池和index池的pg大小
	flushTemplate.Pools = storage.CalculatePg(poolPgs, indexPg)

	if indexCount == 0 {
		flushTemplate.Pools.Sp.Index = 64
		flushTemplate.NoSSD = true
	}

	//副本数量
	flushTemplate.Pools.Sp.Replicas = devs.Replicas
	//磁盘设备列表
	flushTemplate.Device = devsmap
	//主机列表
	flushTemplate.Host = hosts
	//存储网络
	flushTemplate.PublicNetwork = devs.PublicNetwork
	flushTemplate.ClusterNetwork = devs.ClusterNetwork

	var inventoryBuff bytes.Buffer
	inventoryTmpl := template.Must(template.New("host-inventory.tpl").Funcs(funcMap).Parse(string(static.HOST_INVENTORY_TPL)))

	err := inventoryTmpl.Execute(&inventoryBuff, flushTemplate)
	inventoryContent, _ := ioutil.ReadAll(&inventoryBuff)

	var allYMLBuff bytes.Buffer
	allYMLTmpl := template.Must(template.New("all.yml.tpl").Funcs(funcMap).Parse(string(static.ALL_YML_TPL)))

	err = allYMLTmpl.Execute(&allYMLBuff, flushTemplate)
	if err != nil {
		log.Errorf("Can not flush template. Error: %s", err.Error())
		core.WriteResponse(c, errors.WithCode(code.ErrUnknown, "Unknow"), nil)
		return
	}
	allYMLContent, _ := ioutil.ReadAll(&allYMLBuff)

	utils.SaveFile(constand.CEPH_PACKAGE_CONF_PATH+"/host-inventory", string(inventoryContent))
	utils.SaveFile(constand.CEPH_PACKAGE_CONF_PATH+"/group_vars/all.yml", string(allYMLContent))

	// TODO... Need Fan in
	if utils.IsExist("/tmp/ceph_install.log") {
		_ = os.Remove("/tmp/ceph_install.log")
	}
	cmd := exec.Command("bash", "-c", "cd "+constand.CEPH_PACKAGE_BASE_PATH+"  && ./ceph_deploy.sh")

	logfd, _ := os.OpenFile("/tmp/ceph_install.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logfd.Close()

	cmd.Stdout = io.MultiWriter(logfd)
	cmd.Stderr = cmd.Stdout
	go func() {
		err = cmd.Run()
		if err != nil {
			errChan <- err
			return
		}
		doneChan <- struct{}{}
	}()

	select {
	case err := <-errChan:
		db.Client().Table("statuses").Where("id = ?", 1).Update("status", 0)
		core.WriteResponse(c, errors.WithCode(code.ErrCephInstall, "error in ceph install. More detail at /tmp/ceph_install.log. Error: %s", err.Error()), nil)
		return
	case <-doneChan:
		core.WriteResponse(c, nil, map[string]string{"successes": "ok"})
		return
	}
}
