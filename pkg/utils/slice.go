package utils

import "strings"

func RemoveRepeatElement(data []string) []string {
	if len(data) == 0 {
		return []string{}
	}
	var out = make([]string, 0, len(data))
	temp := map[string]struct{}{}
	for _, item := range data {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			out = append(out, item)
		}
	}

	return out
}

func ToString(data []string, sep string) string {
	var out = ""
	for _, item := range data {
		out = out + item + sep
	}
	return strings.TrimRight(out, sep)
}
