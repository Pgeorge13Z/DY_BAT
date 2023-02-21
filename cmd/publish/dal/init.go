package dal

import (
	constants "DY_BAT/pkg/consts"
	sqlscript "DY_BAT/sql/script"
)

func Init() {
	sqlscript.InitDB(constants.MySQLDefaultDSN)
	sqlscript.GetDB().AutoMigrate(&Video{})
}
