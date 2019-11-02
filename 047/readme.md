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

对于当前 http.Transport 而言，**MaxIdleConns** 针对空闲总数，**MaxIdleConnsPerHost** 针对每个服务的空闲连接总数。
- 判别不同服务的方式：网络地址，网络协议，代理。
----