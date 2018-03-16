
package main

import (
	"os"
	)







type inodeUse struct {
    mountPoint string
    level string
}

func (i inodeUse) start() string {

	// mkdir .chaosD

	CreateDirIfNotExist(i.mountPoint + ".chaosD")

        // calculate no of files required


	

        // write files



    return "started latency spike"
}

func (i inodeUse) end() string {

	// remove .chaosD 
	os.RemoveAll(i.mountPoint + ".chaosD")


    return "end iNode usage"
}

