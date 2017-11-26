package playSound

import (
	"os/exec"
)

// Run play uses SOX, install it
// with the following guide http://arielvb.readthedocs.io/en/latest/docs/commandline/sox.html
func Run() {
	exec.Command("play", "result.ogg").Run()
}
