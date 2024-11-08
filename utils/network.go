package utils

import "os/exec"

func NetWorkStatus() bool {
	cmd := exec.Command("ping", "www.baidu.com", "-c", "4", "-W", "5")

	err := cmd.Run()
	if err != nil {
		return false
	}

	return true
}
