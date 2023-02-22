package dal

import (
	"DY_BAT/cmd/publish/kitex_gen/publish"
	constants "DY_BAT/pkg/consts"
	sqlscript "DY_BAT/sql/script"
)

func Init() {
	sqlscript.InitDB(constants.MySQLDefaultDSN)
	sqlscript.GetDB().AutoMigrate(&publish.Video{})
}
