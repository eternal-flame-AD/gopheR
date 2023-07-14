package rcflags

import (
	"log"
	"os/exec"
	"strings"
)

func QueryRFlags() (cflags []string, ldflags []string) {
	cmd := exec.Command("R", "CMD", "config", "--cppflags")
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("failed to run R CMD config --cppflags: %v", err)
	}
	cflags = strings.Fields(string(out))
	cmd = exec.Command("R", "CMD", "config", "--ldflags")
	out, err = cmd.Output()
	if err != nil {
		log.Fatalf("failed to run R CMD config --ldflags: %v", err)
	}
	ldflagsDirty := strings.Fields(string(out))
	for _, flag := range ldflagsDirty {
		if !(strings.HasPrefix(flag, "-Wl,") || strings.HasPrefix(flag, "-f")) {
			ldflags = append(ldflags, flag)
		}
	}
	return
}
