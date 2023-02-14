package main

import (
	follow "DY_BAT/cmd/follow/kitex_gen/follow/followservice"
	"log"
)

func main() {
	svr := follow.NewServer(new(FollowServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
