package main

import (
	"DY_BAT/cmd/publish/kitex_gen/publish/publishservice"
	"DY_BAT/pkg/consts"
	"fmt"
	client2 "github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

func main() {

	r, err := etcd.NewEtcdResolver([]string{consts.EtcdAddress})
	if err != nil {
		fmt.Println(err)
	}

	client, err := publishservice.NewClient(consts.PublishServiceName, client2.WithResolver(r))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(client)

	//PublishActionRequest := publish.NewDouyinPublishActionRequest()
	//PublishActionRequest.Data
	//client.PublishAction(context.Background())
}
