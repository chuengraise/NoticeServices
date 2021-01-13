package main

import (
	_ "NoticeServices/boot"
	"NoticeServices/library/version"
	_ "NoticeServices/router"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gogf/gf/frame/g"
)

var (
	BuildVersion = "0.0"
	BuildTime    = ""
	CommitID     = ""
)

func main() {
	version.ShowLogo(BuildVersion, BuildTime, CommitID)
	g.Server().Run()
}
