#### Protobuf语法

message：定义一个请求或相应的消息格式，可以包含多种类型字段

```protobuf
// proto编译之后，Go文件中，每个消息类型对应一个结构体
message StudentRequest {
    // 每个字段的数值型标识符唯一
    string name = 1;
    int32 age = 2;
    int32 height = 3;
    // message可嵌套
    // HelloReply hello = 4;
}
```

repeated 代表可重复，相当于数组

```protobuf
repeated name = 1;
```

服务定义格式

```protobuf
// service serviceName {
//	rpc method (requestMessage) returns (responseMessage) {}
// }
service StudentService {
    // rpc 远程方法调用
    rpc Student (StudentRequest) returns (StudentResponse) {}
}
```
