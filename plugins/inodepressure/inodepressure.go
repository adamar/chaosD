package inodepressure

import (
	"fmt"
	"github.com/adamar/chaosd/common"
	"github.com/adamar/chaosd/plugins"
	"os"
	"strconv"
	"syscall"
)

type Inodepressure struct {
	MountPoint string
	Level      string
}

func init() {
        plugins.Add("Inodepressure", func(config map[string]string) plugins.Plugin {
                return &Inodepressure{MountPoint: config["MountPoint"], Level: config["Level"]}
        })
}

func (i *Inodepressure) Description() string {
        return "Use up a large no. of inodes"
}

func (i *Inodepressure) Start() string {

	// mkdir .chaosD
	common.CreateDirIfNotExist(i.MountPoint + ".chaosD")

	// calculate no of files required
	numOfInodesToMake := getNoInodesToUseUp("/", i.Level)

	// range over no. of files to create and make an
	// empty file for each
	for s := uint64(0); s < numOfInodesToMake; s++ {
		suffix := strconv.Itoa(int(s))
		common.CreateEmptyFile(i.MountPoint + ".chaosD/file" + suffix)
	}

	return "started latency spike"
}

func (i *Inodepressure) End() string {

	// remove .chaosD
	fmt.Println(i.MountPoint + ".chaosD")
	os.RemoveAll(i.MountPoint + ".chaosD")

	return "end iNode usage"
}

func getNoInodesToUseUp(MountPoint string, num int) uint64 {

	//https://linux.die.net/man/2/statvfs

	var stat syscall.Statfs_t
	syscall.Statfs(MountPoint, &stat)

	// No of inodes which would be 90perc
	ninetyPercent := ((uint64(stat.Files) * uint64(num)) / uint64(100))

	// subtract the no of inodes currently used from the total number we need to use up
	// this final number will be the number of files to be created
	inodesToMake := ninetyPercent - (stat.Files - stat.Ffree)

	return inodesToMake

}
