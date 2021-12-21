# 03 gRPC 的通信模式
## 一元 RPC
> 当客户端调用服务器端的远程方法时，客户端发送请求至服务器端并获得一个响应，与响应一起发送的 还有`状态细节`以及 `trailer` 元数据

```protobuf
syntax = "proto3";
import "google/protobuf/wrappers.proto";
package ecommerce;

service OrderManagement {
  rpc getORder(google.protobuf.StringValue) returns (Order);
}

message Order {
  string id = 1;
  repeated string items = 2;
  string description = 3;
  float price = 4;
  string destination = 5;
}
```

## 服务器端流 RPC
> 在服务器端流 RPC 模式中，服务器端在接收到 客户端的请求消息后，会发回一个响应的序列。这种多个响应所组成的序列也被称为“流”。
> 在将所有的服务器端响应发送完毕之后，服务器端 会以 trailer 元数据的形式将其状态发送给客户端，从而标记流的结束

```protobuf
syntax = "proto3";
import "google/protobuf/wrappers.proto";
package ecommerce;
service OrderManagement {
  rpc searchOrders(google.protobuf.StringValue) returns (stream Order);
}
message Order {
  string id = 1;
  repeated string items = 2;
  string description = 3;
  float price = 4;
  string destination = 5;
}
```

```go
func (s *server)SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {
	for key, order := range orderMap {
		for _, itemStr := range order.Items {
		    if strings.Contains(itemStr, searchQuery.Value){
			    err := stream.Send(&order)
				if err != nil {
				    return fmt.Errorf("error sending message to stream error: %s", err.Error())	
                }
				break
            }   	
        }       
    }
	return nil
}
```

```go
client := pd.NewOrderManagementClient(conn)
searchStream, _ := client.SearchOrders(ctx, &wrapper.StringValue{Value:"Google"})
for {
	searchOrder, err := searchStream.Recv()
	if err == io.EOF {
		break
    }   
}
```

## 客户端流 RPC
> 在客户端流 RPC 模式中，客户端会发送多个请求给服务器端，而不再是单个请求。服务器端则会发送一个响应给客户端。
> 但是，服务器端不一定要等到从客户端接收到所有消息后才发送响应
> 可以在接收到流中的一条消息或几条消息之后就发送响应，也可以在读取完流中的所有消息之后再发送响应。

```protobuf
syntax = "proto3";
import "google/protobuf/wrappers.proto";
package ecommerce;
service OrderManagement {
  rpc updateOrders(stream Order) returns (google.protobuf.StringValue);
}

message Order {
  string id = 1;
  repeated string items = 2;
  string description = 3;
  float price = 4;
  string destination = 5;
}
```
```go
func (s *server)UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) {
	for {
	    order, err := stream.Recv()
		if err == io.EOF {
		    return stream.SendAndClose(
				    &wrapper.StringValue{Value: ""}
				)	
        }
		orderMap[Order.Id] = *order
    }
}
```

```go

```
## 双向流 RPC
> 在双向流 RPC 模式中，客户端以消息流的形式发送请求到服务器端， 服务器端也以消息流的形式进行响应。
> 调用必须由客户端发起，但在此 之后，通信完全基于 gRPC 客户端和服务器端的应用程序逻辑

> 核心理念就是一旦调用RPC方法，那么无论是客户端还是服务器端，都可以在任意时间发送消息。这也包括来自任意一端的流结束标记。
> 
> 









