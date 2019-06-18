package networklatency

import (
	"fmt"
	"github.com/adamar/chaosd/plugins"
	"os"
	"os/exec"
)

type Networklatency struct {
	Level    string
	Duration string
}

func init() {
	plugins.Add("Networklatency", func(config map[string]string) plugins.Plugin {
		return &Networklatency{Level: config["Level"], Duration: config["Duration"]}
	})
}

func (n *Networklatency) Description() string {
	return "Increase the latency for a single network interface"
}

func (n *Networklatency) Start() string {
	cmd := "tc"
	args := []string{"qdisc", "add", "dev", "eth0", "root", "netem", "delay", "97ms"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return "success"
}

func (n Networklatency) End() string {
	cmd := "tc"
	args := []string{"qdisc", "del", "dev", "eth0", "root", "netem"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return "success"
}
