package main

import (
	follow "followservice/kitex_gen/follow/followservice"
	"log"
)

func main() {
	svr := follow.NewServer(new(FollowServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
