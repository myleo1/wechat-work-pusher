package main

import (
	"fmt"
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"wechat-work-pusher/cmd"
)

var (
	version   string
	date      string
	goVersion string
)

func main() {
	info := fmt.Sprintf("***Wechat-Work-Pusher %s***\n***BuildDate %s***\n***%s***\n", version, date, goVersion)
	fmt.Print(info)
	initkit.LoadConfig()
	cmd.Execute()
}
