Go Micro 

默认consul进行服务发现 HTTP协议进行通信  protobuf和json进行编解码

顶层Service  初始化 Service 和 Client 的方法，创建一个 RPC 服务

client 请求服务接口

service 监听服务调用接口

Broker 消息发布和订阅的接口

Codec 传输过程编码和解码接口

Registry 实现服务的注册和发现

Selector 客户端级别的负载均衡

Transport 服务之间通信的接口  