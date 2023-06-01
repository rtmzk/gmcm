package template

import (
	"gmcm/pkg/utils/devices"
	"strings"
)

func NewFuncMap() map[string]any {
	return map[string]any{
		"find": func(src, dst string) bool { return strings.Contains(src, dst) },
		"setDevice": func(data []devices.Devices) string {
			var devs []string
			var out string
			for _, dev := range data {
				if !dev.Cached && dev.Type != "ssd" {
					devs = append([]string{`/dev/` + dev.Name}, devs...)
				}
				if !dev.Cached && dev.Type == "ssd" {
					devs = append(devs, `/dev/`+dev.Name)
				}
			}

			for i := 0; i < len(devs); i++ {
				out = out + `'` + devs[i] + `',`
			}
			o := strings.TrimRight(out, ",")

			return o
		},
		"setCacheDevice": func(data []devices.Devices) string {
			var devs []string
			var out = ""
			for _, dev := range data {
				if dev.Cached && dev.Type == "ssd" {
					devs = append(devs, `/dev/`+dev.Name)
				}
			}
			for i := 0; i < len(devs); i++ {
				out = out + `'` + devs[i] + `',`
			}
			o := strings.TrimRight(out, ",")

			return o
		},
	}
}
