package main

import (
	"DY_BAT/pkg/tools"
	"fmt"
	"github.com/godruoyi/go-snowflake"
	"strconv"
	"strings"
)

func main() {

	SnowflakeId := snowflake.ID()
	filename := strings.Join([]string{strconv.Itoa(int(SnowflakeId)), ".mp4"}, "")

	fmt.Println(strings.Join([]string{tools.GetPath(), "/public/", filename}, ""))
}
