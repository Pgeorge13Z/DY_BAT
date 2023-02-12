package dal

import "DY_BAT/cmd/user/dal/db_mysql"

func Init() {
	db_mysql.MySQLInit()
}
