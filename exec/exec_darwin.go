package exec

import (
	"os"
	"os/exec"
	"os/signal"

	"github.com/bitxx/pholcus/config"

	"github.com/bitxx/pholcus/cmd" // cmd版
	"github.com/bitxx/pholcus/web" // web版
)

func run(which string) {
	exec.Command("/bin/sh", "-c", "title", config.FULL_NAME).Start()
	// 选择运行界面
	switch which {
	case "cmd":
		cmd.Run()
	case "web":
		fallthrough
	default:
		ctrl := make(chan os.Signal, 1)
		signal.Notify(ctrl, os.Interrupt, os.Kill)
		go web.Run()
		<-ctrl
	}
}
