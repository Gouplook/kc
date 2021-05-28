module git.900sui.cn/kc/rpcCards

go 1.12

require (
	git.900sui.cn/kc/base v1.0.34
	git.900sui.cn/kc/kcgin v1.0.9
	git.900sui.cn/kc/mapstructure v1.1.5
	git.900sui.cn/kc/redis v1.0.2
	git.900sui.cn/kc/rpcinterface v0.0.0-20210524060910-cedb7d27b349
	github.com/gin-gonic/gin v1.6.2 // indirect
	github.com/golang/protobuf v1.4.0 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/opentracing/opentracing-go v1.1.0
	github.com/shopspring/decimal v1.2.0
	github.com/smallnest/rpcx v0.0.0-20200322104434-654544af007f
	github.com/wendal/errors v0.0.0-20130201093226-f66c77a7882b
	golang.org/x/mod v0.4.1 // indirect
	golang.org/x/tools v0.1.0 // indirect
)

replace git.900sui.cn/kc/rpcinterface => ./rpcinterface
