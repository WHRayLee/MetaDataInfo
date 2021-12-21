# MetaDataInfo

- 配置包
    ```bash
    github.com/spf13/viper
    ```

- 日志包
    ```bash
    go.uber.org/zap
    ```
- SQL解析
  ```bash
  github.com/golang/tools/tree/master/cmd/goyacc
  ```

## Redis
  
```bash
github.com/go-redis/redis
```

## MySQL

```bash
github.com/go-sql-driver/mysql
```
```sql
root:111111@tcp(127.0.0.1:3306)/test?charset=utf8
```

```bash
gorm.io/driver/mysql
get gorm.io/gorm
```
```sql
"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```

## HBase
```bash
github.com/tsuna/gohbase
```
### Create a client
```go
client := gohbase.NewClient("localhost")
```

### Insert a cell
```go
values := map[string]map[string][]byte{"cf": map[string][]byte{"a": []byte{0}}}
putRequest, err := hrpc.NewPutStr(context.Background(), "table", "key", values)
rsp, err := client.Put(putRequest)
```

### Get an entire row
```go
getRequest, err := hrpc.NewGetStr(context.Background(), "table", "row")
getRsp, err := client.Get(getRequest)
```

### Get a specific cell
```go
family := map[string][]string{"cf": []string{"a"}}
getRequest, err := hrpc.NewGetStr(context.Background(), "table", "15",
    hrpc.Families(family))
getRsp, err := client.Get(getRequest)
```

### Get a specific cell with a filter
```go
pFilter := filter.NewKeyOnlyFilter(true)
family := map[string][]string{"cf": []string{"a"}}
getRequest, err := hrpc.NewGetStr(context.Background(), "table", "15",
    hrpc.Families(family), hrpc.Filters(pFilter))
getRsp, err := client.Get(getRequest)
```

### Scan with a filter
```go
pFilter := filter.NewPrefixFilter([]byte("7"))
scanRequest, err := hrpc.NewScanStr(context.Background(), "table",
		hrpc.Filters(pFilter))
scanRsp, err := client.Scan(scanRequest)
```