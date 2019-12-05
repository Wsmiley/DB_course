// rsrc -manifest app.manifest -o rsrc.syso
//go build
package main

import (
	"github.com/lxn/walk"
)

type MyMainWindow struct {
	*walk.MainWindow
}

func main() {
	Select()
}
