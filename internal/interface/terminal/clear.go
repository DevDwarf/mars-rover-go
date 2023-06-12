package terminal

import (
	"os"
	"os/exec"
	"runtime"
)

func Clear() {
	cmd := getClearCommand()
	clearCmd := exec.Command(cmd[0], cmd[1:]...)
	clearCmd.Stdout = os.Stdout
	_ = clearCmd.Run()
}

func getClearCommand() []string {
	switch runtime.GOOS {
	case "windows":
		return []string{"cmd", "/c", "cls"}
	default:
		return []string{"clear"}
	}
}
