package consts

const (
	//mysql
	MySQLDefaultDSN = "root:123456@tcp(localhost:9910)/douyin?charset=utf8&parseTime=True&loc=Local"

	//etcd
	EtcdAddress = "127.0.0.1:2379"

	//user
	UserServiceName = "user_service"
	UserServicePort = 9001

	//comment
	ValidComment   = 1 //评论状态：有效
	InvalidComment = 2 //评论状态：取消
	DateTime       = "2006-01-02 15:04:05"
)
