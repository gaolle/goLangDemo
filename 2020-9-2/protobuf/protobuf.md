### protobuf 安装

* https://github.com/protocolbuffers/protobuf/releases

  下载protoc

  window版本解压之后将bin目录下的protoc.exe放在$GOPATH/bin

  linux 需要编译

  

  安装go语言protobuf插件

* ```shell
  go get -u github.com/golang/protobuf/protoc-gen-go
  ```

安装完之后确保$GOPATH/bin下存在protoc.exe和protoc-gen-go.exe

将$GOPATH/bin加入环境变量

https://geektutu.com/post/quick-go-protobuf.html

处理结构体的驼峰tag

go get github.com/favadi/protoc-go-inject-tag

protoc-go-inject-tag.exe位于$GOPATH/bin









### grpc-go

```shell
go get -u google.golang.org/grpc
```

