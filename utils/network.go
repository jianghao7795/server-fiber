package utils

import "os/exec"

func NetWorkStatus() bool {
	cmd := exec.Command("curl", "www.baidu.com")

	err := cmd.Run()
	if err != nil {
		return false
	}

	return true
}
