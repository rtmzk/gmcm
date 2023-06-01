package devices

import (
	"encoding/json"
	"fmt"
	"gmcm/pkg/utils/ssh"
	"strings"
)

const allRecord = `lsblk -d -o name,rota,size,type | awk '$1 ~ /sd|vd/ && $4 ~ /disk/ {printf "{\"name\": \"%s\",\"type\": \"%s\", \"size\": \"%s\"},",$1,$2,$3}'`

func GetRemoteDevices(cli *ssh.Client) ([]Devices, error) {
	var devices []Devices

	// 远程获取所有设备列表
	o, err := cli.Output(allRecord)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// 字符串操作,删除最后一个逗号,字符串前后加上[]
	// "{"name": "sda"},{"name":"sdb"}," => "[{"name": "sda"},{"name":"sdb"}]"
	s := strings.TrimSuffix(string(o), ",")
	s = "[" + s + "]"

	err = json.Unmarshal([]byte(s), &devices)
	if err != nil {
		return nil, err
	}

	// fileterSystemDevice
	// 过滤掉有文件系统的设备,如果有文件系统则从切片中删除
	var checkRes []byte
	var k = 0
	for idx, dd := range devices {
		devices[idx].Type = DeviceTypeMapping[dd.Type]
		devices[idx].DefaultType = DeviceTypeMapping[dd.Type]
		devices[idx].Enabled = true

		if dd.Type == "ssd" {
			devices[idx].Cached = true
		} else {
			devices[idx].Cached = false
		}

		p := DeviceDefaultPathPrefix + dd.Name
		command := `lsblk -r ` + p + ` | awk '$0 ~ /part|lvm/'`
		checkRes, err = cli.Output(command)
		if err != nil {
			fmt.Println(err.Error())
		}

		if len(checkRes) == 0 {
			devices[k] = devices[idx]
			k = k + 1
		}
	}

	return devices[:k], nil
}
