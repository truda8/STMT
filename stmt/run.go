package stmt

import (
	"os/exec"
)

// windows 执行命令，后台
func Exec_win_daemon(c string) {
	cmd := exec.Command("powershell", "/C", c)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Start()
	if err == nil {
		cmd.Process.Release()
	}
}

// linux 执行命令，后台
func Exec_linux_daemon(c string) {
	cmd := exec.Command("bash", "-c", c)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Start()
	if err == nil {
		cmd.Process.Release()
	}
}

// windows 执行命令，终端
func Exec_win(c string) {
	command := "wt powershell -NoExit \"" + c + "\""
	cmd := exec.Command("powershell", "/C", command)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Start()
	if err == nil {
		cmd.Process.Release()
	}
}

// linux 执行命令，终端
func Exec_linux(c string) {
	command := "gnome-terminal -- $SHELL -c '" + c + ";exec $SHELL'"
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Start()
	if err == nil {
		cmd.Process.Release()
	}
}
