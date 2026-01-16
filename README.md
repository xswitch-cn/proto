# proto - XSwitch XCC API Protocol Buffer 文件及各语言SDK

本仓库包含以下内容：

- proto：Protocol Buffer描述文件。
- go：Go语言SDK，配合[xctrl](https://github.com/xswitch-cn/xctrl)使用。
- java：Java语言SDK。
- docs：文档。
- xctrl：Go Protobuf代码生成器和插件。

相关文档：

- [xctrl](https://github.com/xswitch-cn/xctrl)：Go语言SDK使用和开发文档。
- 协议参考文档参见：<https://github.com/xswitch-cn/proto/src/branch/main/docs>。
- 示例：<https://git.xswitch.cn/xswitch/xcc-examples/src/branch/master/go> 。
- XCC API文档：<https://docs.xswitch.cn/xcc-api/> 。

## 使用和开发


### 修改proto

1. 修改[proto](proto/)下的[xctrl.proto](proto/xctrl/xctrl.proto)或者[cman.proto](proto/cman/cman.proto)
2. 如果需要实现ChannelEvent则需要在[xctrl.go](xctrl/cmd/protoc-gen-xctrl/plugin/xctrl/xctrl.go)下的**channelMethodTimeout**的Map中添加Key和Value, Key是Rpc**方法名**,Value是延迟**秒数**

### 开发教程

1. 克隆该项目到本地：

```shell
git clone https://github.com/xswitch-cn/xctrl.git
cd xctrl
```

2. Protocol Buffers编译器（protoc）

```shell
brew install protobuf
```

3. 安装protoc-gen-doc依赖：

- 推荐方式：

```shell
make setup  
```

- 手动安装: 

```shell
go install github.com/chuanlinzhang/protoc-gen-doc/cmd/protoc-gen-doc@v0.0.2
```

4. 根据需要生成相应语言的代码:

- 生成Go代码

```shell
make proto
```

- 生成Java代码

```shell
make java
```

**注意**：Java代码仅为示例参考，`.proto`在后续更新时不会同步更新。如果使用Java的请自行执行`make proto`生成最新代码。

- 生成Markdown文档

```shell
make doc-md
```

- 生成HTML文档

```shell
make doc-html
```

## 测试

```shell
go run main.go
make test
```
