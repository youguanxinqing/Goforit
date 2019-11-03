# 知识点

http.Get 允许直接发起 http 请求(需要指明协议)(code: UseHttp)：
```go
resp, err := http.Get("http://www.baidu.com")
if err != nil {
    ...
}
...
```

http.Get 借助 DefaultClient.Get 完成任务。
- DefaultClient 是 *http.Client 指针类型(`DefaultClient = &Client{}`)。可以做到开箱即用：
```go
var httpClient1 http.Client
resp2, err := httpClient1.Get(url1)
```

1. http.Client 中的 Transpost 字段代表什么？

**向网络服务发送 HTTP 请求，并从网络服务接收 HTTP 响应的操作过程。**

该字段的方法 RoundTrip 应该实现单次 HTTP 事务（或者说基于 HTTP 协议的单次交互）需要的所有步骤。

如果没有指定 Transport 的值，会使用默认值 DefaultTransport。

http.Client.Timeout 是 time.Duration 类型，其值可为 0，表示没有设置超时时间。

----
**深入理解 DefaultTransPort**

```go
var DefaultTransport RoundTripper = &Transport{
	Proxy: ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}
```
- IdleConnTimeout 表示 空闲连接在多久之后会被关闭。该值为 0 时，表示不关闭空闲连接。
- ResponseHeaderTimeout 表示 从系统内核收到请求，到内核收到响应的最大时长。DefaultTransport 没有该字段。
- ExpectContinueTimeout 表示 客户端接收请求报文头后，等待接收第一个响应报文头的最长时间。当客户端想把很大的报文 POST 给服务器时，可先发送一个含有 **“Expect: 100-continue”** 的请求报文头，询问服务器是否愿意接收该报文。如果该字段不大于 0，表示请求报文会立即发送出去，可能造成网络资源浪费。
- TLSHandshakeTimeout 表示 建立 TLS 的握手超时时间。该值为 0 表示没有限制。
- KeepAlive 表示 TCP 层的探测包。

对于当前 http.Transport 而言，**MaxIdleConns** 针对空闲总数，**MaxIdleConnsPerHost** 针对每个服务的空闲连接总数。
- 判别不同服务的方式：网络地址，网络协议，代理。
- MaxIdleConnsPerHost 默认值为 2。

----

连接复用的两种情况：
- 针对同一个服务，有新的 HTTP 请求被递交，连接被再次使用。
- 不再对该网络服务器有 HTTP 请求，连接被闲置。(如果分配给一个服务的连接过多，也会造成空闲连接)
- 如果想杜绝空闲连接，则 DisableKeepAlives 置为 true。(每一次 HTTP 请求被提交都产生一个新的连接，明显增重网络服务与客户端压力)

# http.Server 类型的 ListenAndServe 方法都做了哪些事情？ 扩展(code: HttpServer)

ListenAndServe() 会监听一个基于 TCP 的网络地址，并对接收到的 HTTP 请求进行处理。
- 该方法默认开启针对网络连接的存活探测机制，保证连接持久。
- 当被外界关掉时，它会返回一个由 http.ErrServerClosed 变量代表的错误值。

ListenAndServe 的处理流程：
- 检查 http.Server 类型值的 Addr 字段(ip:port)。如果为空，默认用 “:http”
- 调用 net.Listen，在已确定的网络地址上启动基于 TCP 的监听。
- 检查 net.Listen 函数返回的错误值，如果值不为 nil，直接返回。否则调用对应的 Handler 处理请求。

net.Listen 做了哪些工作：
- 解析参数值中包含的网络地址隐含的 IP 和 端口。
- 根据给定的网络协议，确定监听方法，并开始监听。

http.Serve 类型的 Serve 方法如何处理请求？
- 在 for 循环当中，调用 Accept()，等待客户端连接。Accept 返回两个结果，一个 net.Conn，一个 error。
- 如果 err 为 nil，则第一个结果值会被包装成 *http.conn 类型，然后在新的 goroutine 中调用它的 serve 方法，来处理当前的 HTTP 请求。
- 如果 err 不为 nil，循环终止。(如果 err 代表的是暂时性错误，循环不会终止，而是在一段时间之后开始下一次迭代)

# 如果关闭一个 http 服务器 扩展

server.Shutdown()