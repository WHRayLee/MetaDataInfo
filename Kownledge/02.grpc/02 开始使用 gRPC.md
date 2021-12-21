# 开始使用gRPC
![](/Users/yons/Documents/GitHub/MetaDataInfo/Kownledge/02.grpc/asset/ProductInfo服务的客户端–服务器端交互.png)

## 2.1 创建服务定义
> protocol buffers 可以定义消息类型和服务类型
- 消息包含字段，每个字段由其类型和唯一索引值进行定义
- 服务则包含方法，每个方法由其类型、输入参数和输出参数进行定义。

### 2.1.1 定义消息类型
> 消息（message）是客户端和服务器端交换的数据结构

- ProductID
    1. google.protobuf.StringValue
    2. string

```protobuf
message ProductID {
  string value = 1;
}
```

- Product

```protobuf
message Product {
  string id = 1;
  string name = 2;
  string description = 3;
  float price = 4;
}
```

### 2.1.2 定义服务类型
> 按照 protocol buffers的规则，远程方法只能有一个参数，并只能返回一个值。

- addProduct

```protobuf
rpc addProduct(Product) returns (google.protobuf.StringValue);
```

- getProduct

```protobuf
rpc getProduct(google.protobuf.StringValue) returns (Products);
```

```protobuf
syntax = "proto3";
package Product.V1;

import "google/protobuf/wrappers.proto";

service ProductInfo {
  rpc addProduct(Product) returns (ProductID);
  rpc getProduct(ProductID) returns (Product);
}

message Product {
  string id = 1;
  string name = 2;
  string description = 3;
}

message ProductID {
  string value = 1;
}
```
1. 服务定义首先要指定所使用的 protocol buffers 版本
2. 为了避免协议消息类型之间的命名冲突，这里使用了包名

## 2.2 实现
### 2.2.1 开发服务
![基于服务定义的微服务和消费者](/Users/yons/Documents/GitHub/MetaDataInfo/Kownledge/02.grpc/asset/基于服务定义的微服务和消费者.png)

```go
require ( 
	github.com/gofrs/uuid v3.2.0 
	github.com/golang/protobuf v1.3.2 
	github.com/google/uuid v1.1.1 
	google.golang.org/grpc v1.24.0
)
```

1. 使用如下命令安装 gRPC 库
```go
go get -u google.golang.org/grpc
```
2. 使用如下命令安装 Go 语言的 protoc 插件
```go
go get -u github.com/golang/protobuf/protoc-gen-go
```
3. 编译protobuf文件
```bash
protoc -I ecommerce \ 
ecommerce/product_info.proto \  
--go_out=plugins=grpc:<module_dir_path>/ecommerce
```
    1. 指定源 proto 文件和依赖的 proto 文件的目录路径（通过 --proto_path 或 -I 命令行标记来指定）。如果不指定该值，则将使用`当前目录`作为源目录。在这个目录下，需要根据包名来存放依赖的proto文件
    2. 指定希望编译的 proto 文件路径。编译器将阅读该文件并生成输出的 Go 文件
    3. 指定生成的代码要存放的目标目录

