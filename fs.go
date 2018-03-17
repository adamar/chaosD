
package main

import (
	"os"
	"strconv"
	"syscall"
	"fmt"
	)







type inodeUse struct {
    mountPoint string
    level int
    duration int
}

func (i inodeUse) start() string {

	// mkdir .chaosD
	CreateDirIfNotExist(i.mountPoint + ".chaosD")

        // calculate no of files required
	numOfInodesToMake := getNoInodesToUseUp("/", i.level)

	// range over no. of files to create and make an 
	// empty file for each
	for s := uint64(0); s < numOfInodesToMake; s++ {
		suffix := strconv.Itoa(int(s))
		createEmptyFile(i.mountPoint + ".chaosD/file" + suffix)
	}




    return "started latency spike"
}

func (i inodeUse) end() string {

	// remove .chaosD 
	fmt.Println(i.mountPoint + ".chaosD")
	os.RemoveAll(i.mountPoint + ".chaosD")


    return "end iNode usage"
}



func getNoInodesToUseUp(mountPoint string, num int) uint64 {

        //https://linux.die.net/man/2/statvfs

        var stat syscall.Statfs_t
        syscall.Statfs(mountPoint, &stat)

        // No of inodes which would be 90perc
        ninetyPercent := ( ( uint64(stat.Files) * uint64( num ) ) / uint64(100) )

        // subtract the no of inodes currently used from the total number we need to use up
        // this final number will be the number of files to be created
        inodesToMake := ninetyPercent - ( stat.Files - stat.Ffree  )

        return inodesToMake

}











