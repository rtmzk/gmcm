package utils

import "os/user"

func UserHome() string {
	u, err := user.Current()
	if err != nil {
		return "/root"
	}
	return u.HomeDir
}
