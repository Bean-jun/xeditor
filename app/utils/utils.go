package utils

import (
	"bytes"
	"os/exec"
	"regexp"
	"runtime"
	"syscall"
)

func RefindNumber(strs string) string {
	patten := regexp.MustCompile(`cpython-(\d+\.\d+\.\d+)-windows-x86_64-none`)
	match := patten.FindStringSubmatch(strs)
	// log.Println(match)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func ExecuteCmdNoWindow(execute string, args ...string) (string, string, error) {
	var b bytes.Buffer
	var c bytes.Buffer
	cmd := exec.Command(execute, args...)

	if runtime.GOOS == "windows" {
		// 修复window下黑窗口弹出问题
		cmd.SysProcAttr = &syscall.SysProcAttr{
			// HideWindow:    true,
			CreationFlags: 0x08000000,
		}
	}

	cmd.Stdout = &b
	cmd.Stderr = &c
	err := cmd.Run()
	if err != nil {
		return "", "", err
	}
	return b.String(), c.String(), nil
}
