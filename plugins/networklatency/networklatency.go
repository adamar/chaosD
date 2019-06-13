package networklatency

import (
	"fmt"
	"github.com/adamar/chaosd/plugin"
	"os"
	"os/exec"
)

type Networklatency struct {
}

func init() {
	fmt.Println("Registering networklatency")
	plugin.RegisterPlugin(&Networklatency{})
}

func EnableLatency() {
	cmd := "tc"
	args := []string{"qdisc", "add", "dev", "eth0", "root", "netem", "delay", "97ms"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func DisableLatency() {
	cmd := "tc"
	args := []string{"qdisc", "del", "dev", "eth0", "root", "netem"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
