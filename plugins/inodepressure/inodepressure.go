
package inodepressure

import (
	"fmt"
	"github.com/adamar/chaosd/plugin"
)


type Inodepressure struct {
}

func init() {
	fmt.Println("Registering inodepressure")
	plugin.RegisterPlugin(&Inodepressure{})
}

